package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Registry struct {
	Isbn           string         `db:"isbn"`
	Title          string         `db:"title"`
	Author         string         `db:"author"`
	Translator     string         `db:"translator"`
	PrintAmount    int            `db:"print_amount"`
	SelfPublished  bool           `db:"self_published"`
	Partner        string         `db:"partner"`
	ConfirmationId string         `db:"confirmation_id"`
	Created        types.DateTime `db:"created"`
	Updated        types.DateTime `db:"updated"`
}

func (c Registry) TableName() string {
	return "registries"
}

func FindRegistryByConfirmationId(db dbx.Builder, id string) (*Registry, error) {
	registry := &Registry{}

	err := db.Select().Where(dbx.HashExp{"confirmation_id": id}).One(&registry)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return registry, err
}

func AddRegistry(db dbx.Builder, registry *Registry) error {
	if err := db.Model(registry).Insert(); err != nil {
		return err
	}

	return nil
}
