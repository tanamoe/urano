package jobs

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tanamoe/urano/models"
)

func GetLatestRegistries(app *pocketbase.PocketBase) {
	lastPage, err := getLastPage(app)

	if err != nil {
		return
	}

	for i := 1; i <= lastPage; i++ {
		var registries []models.Registry

		err := getTable(app, i, func(r []models.Registry) {
			registries = append(registries, r...)
		})

		if err != nil {
			app.Logger().Error("Cannot get table from PPDVN", "error", err)
			return
		}

		if err := insert(app, registries); err != nil {
			break
		}
	}
}

func insert(app *pocketbase.PocketBase, registries []models.Registry) error {
	for _, registry := range registries {
		if _, err := models.FindRegistryByConfirmationId(app.DB(), registry.ConfirmationId); err == nil {
			app.Logger().Info("Something already exists in database, stopping job...")
			return nil
		}

		if err := models.AddRegistry(app.DB(), &registry); err != nil {
			app.Logger().Error("An error occurred while inserting", "error", err)
			return err
		}
	}

	return nil
}

func getTable(app *pocketbase.PocketBase, page int, invoke func(registries []models.Registry)) error {
	c := colly.NewCollector()

	c.OnHTML("div#list_data_return tbody", func(e *colly.HTMLElement) {
		results, err := parseTable(e)

		if err != nil {
			return
		}

		invoke(results)
	})

	c.OnRequest(func(r *colly.Request) {
		app.Logger().Info("Getting paged registries", "url", r.URL)
	})

	c.Visit(fmt.Sprintf("https://ppdvn.gov.vn/web/guest/ke-hoach-xuat-ban?p=%d", page))

	return nil
}

func parseTable(e *colly.HTMLElement) ([]models.Registry, error) {
	var registries []models.Registry

	e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
		print_amount_string := e.ChildText("td:nth-child(6)")

		i, err := strconv.Atoi(print_amount_string)
		if err != nil {
			return
		}

		registry := models.Registry{
			Isbn:           e.ChildText("td:nth-child(2)"),
			Title:          e.ChildText("td:nth-child(3)"),
			Author:         e.ChildText("td:nth-child(4)"),
			Translator:     e.ChildText("td:nth-child(5)"),
			PrintAmount:    i,
			SelfPublished:  e.ChildText("td:nth-child(7)") == "x",
			Partner:        e.ChildText("td:nth-child(8)"),
			ConfirmationId: e.ChildText("td:nth-child(9)"),
			Created:        types.NowDateTime(),
			Updated:        types.NowDateTime(),
		}

		registries = append(registries, registry)
	})

	return registries, nil
}

func getLastPage(app *pocketbase.PocketBase) (int, error) {
	var page int

	c := colly.NewCollector()

	c.OnHTML(".pagination ul li:last-child a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")

		r := regexp.MustCompile("\\d+$")

		match := r.FindString(href)

		page_value, err := strconv.Atoi(match)

		if err != nil {
			return
		}

		page = page_value
	})

	c.OnRequest(func(r *colly.Request) {
		app.Logger().Info("Getting registries", "url", r.URL)
	})

	c.Visit("https://ppdvn.gov.vn/web/guest/ke-hoach-xuat-ban")

	return page, nil
}
