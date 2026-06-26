package mapper

import (
	api "buf.build/gen/go/tanamoe/urano/protocolbuffers/go/urano/api/v1beta1"
	"github.com/tanamoe/urano/internal/models"
	"google.golang.org/genproto/googleapis/type/date"
)

func NewRegistryResponse(registry models.Registry) *api.Registry {
	r := &api.Registry{
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
		r.Isbn = &registry.Isbn.String
	}

	if registry.Author.Valid {
		r.Author = &registry.Author.String
	}

	if registry.Translator.Valid {
		r.Translator = &registry.Translator.String
	}

	if registry.PrintAmount.Valid {
		r.PrintAmount = &registry.PrintAmount.Int32
	}

	if registry.SelfPublish.Valid {
		r.SelfPublish = &registry.SelfPublish.Bool
	}

	if registry.Partner.Valid {
		r.Partner = &registry.Partner.String
	}

	return r
}
