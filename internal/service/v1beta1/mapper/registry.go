package mapper

import (
	api "buf.build/gen/go/tanamoe/urano/protocolbuffers/go/urano/api/v1beta1"
	"github.com/tanamoe/urano/internal/models"
	"google.golang.org/genproto/googleapis/type/date"
)

func NewRegistryResponse(registry models.Registry) *api.Registry {
	builder := &api.Registry_builder{
		Id:             registry.ID.String(),
		RegistrationId: registry.RegistrationID,
		Title:          registry.Title,
		RegistrationDate: &date.Date{
			Year:  int32(registry.RegistrationDate.Time.Year()),
			Month: int32(registry.RegistrationDate.Time.Month()),
			Day:   int32(registry.RegistrationDate.Time.Day()),
		},
	}

	if registry.Isbn.Valid {
		builder.Isbn = &registry.Isbn.String
	}

	if registry.Author.Valid {
		builder.Author = &registry.Author.String
	}

	if registry.Translator.Valid {
		builder.Translator = &registry.Translator.String
	}

	if registry.PrintAmount.Valid {
		builder.PrintAmount = &registry.PrintAmount.Int32
	}

	if registry.SelfPublish.Valid {
		builder.SelfPublish = &registry.SelfPublish.Bool
	}

	if registry.Partner.Valid {
		builder.Partner = &registry.Partner.String
	}

	return builder.Build()
}
