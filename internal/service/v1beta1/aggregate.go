package v1beta1

import (
	"context"

	"buf.build/gen/go/tanamoe/urano/connectrpc/go/urano/v1beta1/uranov1beta1connect"
	uranov1beta1 "buf.build/gen/go/tanamoe/urano/protocolbuffers/go/urano/v1beta1"
	"connectrpc.com/connect"
	"github.com/tanamoe/urano/providers/fahasa"
)

var _ uranov1beta1connect.AggregateServiceHandler = (*AggregateServer)(nil)

type AggregateServer struct {
	uranov1beta1connect.UnimplementedAggregateServiceHandler

	fahasaClient fahasa.Client
}

func (s *AggregateServer) GetFahasaProduct(
	ctx context.Context,
	req *connect.Request[uranov1beta1.GetFahasaProductRequest],
) (*connect.Response[uranov1beta1.GetFahasaProductResponse], error) {
	product, err := s.fahasaClient.Product(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := &uranov1beta1.GetFahasaProductResponse{Product: &uranov1beta1.AggregateProduct{
		Price: product.Price,
	}}
	return connect.NewResponse(res), nil
}
