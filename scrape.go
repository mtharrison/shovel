package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapePaginated(manifest Manifest, next func(int, []map[string]string) (bool, string)) []map[string]string {

	url := manifest.URL

	var results []map[string]string
	i := 0

	for {
		res := ScrapePage(url, manifest)
		keep, nextUrl := next(i, res)

		if keep {
			results = append(results, res...)
		}

		if nextUrl == "" {
			break
		}

		i++
		url = nextUrl
	}

	return results
}

func ScrapePage(url string, manifest Manifest) []map[string]string {

	var results []map[string]string

	c := colly.NewCollector()

	c.OnHTML(manifest.SingleItemSelector, func(e *colly.HTMLElement) {

		m := make(map[string]string, len(manifest.AttributeSelectors))

		for key, selector := range manifest.AttributeSelectors {
			val := strings.TrimSpace(e.ChildText(selector))
			if val == "" {
				return
			}
			m[key] = val
		}

		results = append(results, m)
	})

	fmt.Println("Visiting", url)

	c.Visit(url)

	return results
}
