package v1beta1

import (
	"context"
	"log/slog"
	"strconv"
	"strings"

	apiconnect "buf.build/gen/go/tanamoe/urano/connectrpc/go/urano/api/v1beta1/apiv1beta1connect"
	api "buf.build/gen/go/tanamoe/urano/protocolbuffers/go/urano/api/v1beta1"
	types "buf.build/gen/go/tanamoe/urano/protocolbuffers/go/urano/types/v1beta1"
	"connectrpc.com/connect"
	"github.com/tanamoe/urano/providers/fahasa"
)

var _ apiconnect.AggregateServiceHandler = (*aggregate)(nil)

type aggregate struct {
	apiconnect.UnimplementedAggregateServiceHandler

	fahasaClient fahasa.Client
}

func NewAggregateServer() apiconnect.AggregateServiceHandler {
	fahasaClient := fahasa.NewClient()
	return &aggregate{
		fahasaClient: fahasaClient,
	}
}

func (s *aggregate) GetFahasaProduct(
	ctx context.Context,
	req *connect.Request[api.GetFahasaProductRequest],
) (*connect.Response[api.GetFahasaProductResponse], error) {
	product, err := s.fahasaClient.Product(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	x, y, z, err := normaliseSize(product.Size)
	if err != nil {
		return nil, err
	}

	images := make([]string, 0, len(product.MediaGallery.Images))
	for _, image := range product.MediaGallery.Images {
		images = append(images, image.File)
	}

	res := &api.GetFahasaProductResponse{Product: &api.AggregateProduct{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Sku:         product.SKU,
		Weight:      float64(product.Weight),
		Size: &types.Size{
			X: x,
			Y: y,
			Z: z,
		},
		PageCount: uint32(product.PageCount),
		Images:    images,
	}}
	return connect.NewResponse(res), nil
}

func (s *aggregate) ListFahasaProduct(
	ctx context.Context,
	req *connect.Request[api.ListFahasaProductRequest],
) (*connect.Response[api.ListFahasaProductResponse], error) {
	product, err := s.fahasaClient.ListByCategory(ctx, fahasa.ListByCategoryParams{
		CategoryID: req.Msg.CategoryId,
		PageSize:   12,
		Page:       1,
	})
	if err != nil {
		return nil, err
	}

	products := make([]*api.AggregateProduct, 0, len(product.ProductList))
	for _, product := range product.ProductList {
		price, _ := normalisePrice(product.ProductPrice)
		products = append(products, &api.AggregateProduct{
			Name:   product.ProductName,
			Price:  price,
			Images: []string{product.ImageSrc},
		})
	}

	res := &api.ListFahasaProductResponse{Products: products}
	return connect.NewResponse(res), nil
}

func normalisePrice(s string) (int64, error) {
	normalised := strings.ReplaceAll(s, ".", "")
	return strconv.ParseInt(normalised, 10, 64)
}

func normaliseSize(s string) (x, y, z float64, err error) {
	sizeNum := strings.TrimSuffix(s, "cm")
	sizeSeparator := strings.ReplaceAll(sizeNum, ",", ".")
	sizeNormalised := strings.ReplaceAll(sizeSeparator, " ", "")
	sizes := strings.Split(sizeNormalised, "x")

	switch len(sizes) {
	case 3:
		if z, err = strconv.ParseFloat(sizes[2], 64); err != nil {
			slog.Error("cannot parse z", "error", err)
			return
		}
		fallthrough
	case 2:
		if y, err = strconv.ParseFloat(sizes[1], 64); err != nil {
			slog.Error("cannot parse y", "error", err)
			return
		}
		fallthrough
	case 1:
		if x, err = strconv.ParseFloat(sizes[0], 64); err != nil {
			slog.Error("cannot parse x", "error", err)
			return
		}
	}

	return
}
