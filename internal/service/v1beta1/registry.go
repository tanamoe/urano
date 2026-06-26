package v1beta1

import (
	"context"

	apiconnect "buf.build/gen/go/tanamoe/urano/connectrpc/go/urano/api/v1beta1/apiv1beta1connect"
	api "buf.build/gen/go/tanamoe/urano/protocolbuffers/go/urano/api/v1beta1"
	"connectrpc.com/connect"
	"github.com/tanamoe/urano/internal/models"
	"github.com/tanamoe/urano/internal/service/v1beta1/mapper"
)

var _ apiconnect.AggregateServiceHandler = (*aggregate)(nil)

type registry struct {
	repo *models.Queries
}

func NewRegistryServer(repo *models.Queries) apiconnect.RegistryServiceHandler {
	return &registry{repo: repo}
}

func (s *registry) ListRegistry(
	ctx context.Context,
	req *connect.Request[api.ListRegistryRequest],
) (*connect.Response[api.ListRegistryResponse], error) {
	registries, err := s.repo.ListRegistry(ctx, models.ListRegistryParams{
		Limit:  req.Msg.PageSize,
		Offset: req.Msg.Skip,
	})
	if err != nil {
		return nil, err
	}

	responses := make([]*api.Registry, 0, len(registries))
	for _, registry := range registries {
		responses = append(responses, mapper.NewRegistryResponse(registry))
	}

	return connect.NewResponse(&api.ListRegistryResponse{
		Registries: responses,
	}), nil
}
