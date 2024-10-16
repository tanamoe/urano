package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"tana.moe/urano/apis"
	"tana.moe/urano/hooks"
	_ "tana.moe/urano/migrations"
	"tana.moe/urano/services"
)

func main() {
	err := godotenv.Load()
	if (err != nil) && (!errors.Is(err, os.ErrNotExist)) {
		log.Fatal("Error loading .env file", err)
		return
	}

	app := pocketbase.New()

	if err := registerMiddleware(app); err != nil {
		log.Fatal(err)
		return
	}

	if err := registerApis(app); err != nil {
		log.Fatal(err)
		return
	}

	if err := hooks.RegisterHooks(app); err != nil {
		log.Fatal(err)
		return
	}

	if err := startServices(app); err != nil {
		log.Fatal(err)
		return
	}

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: false,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
		return
	}
}

func registerMiddleware(
	app *pocketbase.PocketBase,
) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return nil
	})
	return nil
}

func registerApis(
	app *pocketbase.PocketBase,
) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		apis.RegisterApis(app, e)
		return nil
	})
	return nil
}

func startServices(
	app *pocketbase.PocketBase,
) error {
	if err := services.Start(app); err != nil {
		return err
	}
	return nil
}
