package main

import (
    "log"
    "path"

    "github.com/gocolly/colly/v2"

    . "github.com/phdah/advanced-search/internal"
)

func main() {
    esAdress := "http://localhost:9200"
    esUser := "elastic"
    esPass := "hackathon123"

    esIndex := "avinode_api"

    es := Es(esAdress, esUser, esPass)

    // Instantiate default collector
    c := colly.NewCollector()

    // Start scraping the website
    URLS := []string{
        "https://developer.avinodegroup.com/docs/introduction",
        "https://developer.avinodegroup.com/docs/api-basics",
        "https://developer.avinodegroup.com/docs/error-handling-guide",
        "https://developer.avinodegroup.com/docs/sandbox",
        "https://developer.avinodegroup.com/docs/terminology",
        "https://developer.avinodegroup.com/docs/getting-started-webhooks",
        "https://developer.avinodegroup.com/docs/avinode-webhooks",
        "https://developer.avinodegroup.com/docs/schedaero-webhooks",
        "https://developer.avinodegroup.com/docs/working-with-deep-links",
        "https://developer.avinodegroup.com/docs/brand-guidelines",
    }

    // Iterate over URLs
    for _, url := range URLS {
        documentName := ""
        document := ""

        // On request log the URL
        c.OnRequest(func(r *colly.Request) {
            documentName = path.Base(r.URL.String())
        })

        // Get all paragraphs
        c.OnHTML("h1, h2, h3, h4, h5, h6, p, li", func(e *colly.HTMLElement) {
            paragraph := e.Text
            document += " " + paragraph
        })

        // Visit the page
        err := c.Visit(url)
        if err != nil {
            log.Fatal(err)
        }

        // Ensure everything is processed before indexing
        c.Wait()

        // Index the document
        es.Put(esIndex, documentName, document)
    }
}
