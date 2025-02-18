package main

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"github.com/tanamoe/urano/jobs"
	_ "github.com/tanamoe/urano/migrations"
)

func main() {
	app := pocketbase.New()

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	app.Cron().MustAdd("getLatestRegistries", "0 0 1 1 0", func() {
		jobs.GetLatestRegistries(app)
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
