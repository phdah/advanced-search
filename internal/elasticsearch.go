package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// ESClient represents your Elasticsearch client
type ESClient struct {
	*elasticsearch.Client
}

// Source represents the _source object in each hit
type Source struct {
	Content string `json:"content"`
}

// Hit represents a single hit in the Elasticsearch response
type Hit struct {
	Source Source `json:"_source"`
}

// Hits represents the hits section in the Elasticsearch response
type Hits struct {
	Hits []Hit `json:"hits"`
}

// Response represents the top-level Elasticsearch response structure
type ESResponse struct {
	Hits Hits `json:"hits"`
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

// Parse extracts all "content" fields, concatenates them, and returns a single string
func (client *ESClient) Parse(response *esapi.Response) (string, error) {
	if response == nil {
		return "", io.EOF
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the JSON into the ESResponse struct
	var esResponse ESResponse
	if err := json.Unmarshal(body, &esResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	// Use a StringBuilder for efficient string concatenation
	var sb strings.Builder
	for _, hit := range esResponse.Hits.Hits {
		sb.WriteString(hit.Source.Content)
		sb.WriteString(" ") // Add a space or any separator between content strings, if needed
	}

	// Convert the StringBuilder to a string and return
	return sb.String(), nil
}
