package internal

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ESClient struct {
	*elasticsearch.Client
}

func Es(adress string, user string, pass string) *ESClient {
	// Elasticsearch client configuration
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			adress,
		},
		Username: user,
		Password: pass,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return &ESClient{Client: es}
}

func (Es *ESClient) Put(index string, id string, body string) *esapi.Response {
	// Prepare index request
	jsonBody := fmt.Sprintf(`{"content": %q}`, body)
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(jsonBody),
		Refresh:    "true",
	}

	// Perform the request
	res, err := req.Do(context.Background(), Es)
	if err != nil {
		log.Printf("Error indexing id %s: %s", id, err)
		return nil
	}

	// Close response body only if res is not nil
	if res != nil {
		defer res.Body.Close()
		if res.IsError() {
			log.Printf("Error response from Elasticsearch for id %s: %s", id, res.String())
		} else {
			log.Printf("Successfully indexed id %s\n", id)
		}
	}
	return res
}

func (Es *ESClient) Get(index string, id string, body string) string {

	return ""
}
