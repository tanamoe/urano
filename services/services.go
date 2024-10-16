package services

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func Start(app *pocketbase.PocketBase) error {
	if err := startPreMigrationService(app); err != nil {
		return err
	}
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		if err := startPostMigrationService(app); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func startPreMigrationService(app *pocketbase.PocketBase) error {
	return nil
}

func startPostMigrationService(app *pocketbase.PocketBase) error {
	return nil
}
