package internal

import (
	"bytes"
	"context"
	"encoding/json"
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

func (client *ESClient) Get(index string, query string, size int) (*esapi.Response, error) {
	// Construct the request body
	body := map[string]interface{}{
		"size":    size,
		"_source": true,
		"query": map[string]interface{}{
			"match": map[string]string{
				"content": query,
			},
		},
	}

	// Convert the body to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	// Perform the search request
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  &buf,
	}

	res, err := req.Do(context.Background(), client.Client)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %w", err)
	}

	if res.IsError() {
		return res, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	return res, nil
}
