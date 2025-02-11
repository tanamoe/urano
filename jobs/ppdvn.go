package jobs

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type Registration struct {
	isbn, title, author, translator string
	print_amount                    int
	self_published                  bool
	partner, confirmation_id        string
}

func GetTable(page int) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div#list_data_return tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			print_amount_string := e.ChildText("td:nth-child(6)")

			i, err := strconv.Atoi(print_amount_string)
			if err != nil {
				return
			}

			registration := Registration{
				isbn:            e.ChildText("td:nth-child(2)"),
				title:           e.ChildText("td:nth-child(3)"),
				author:          e.ChildText("td:nth-child(4)"),
				translator:      e.ChildText("td:nth-child(5)"),
				print_amount:    i,
				self_published:  false,
				partner:         e.ChildText("td:nth-child(8)"),
				confirmation_id: e.ChildText("td:nth-child(9)"),
			}
			fmt.Println(registration)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	for i := 0; i <= 5; i++ {
		c.Visit(fmt.Sprintf("https://ppdvn.gov.vn/web/guest/ke-hoach-xuat-ban?p=%d", i))
	}
}
