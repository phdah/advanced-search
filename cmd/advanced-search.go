package main

import (
	"fmt"
	"log"

	. "github.com/phdah/advanced-search/internal"
)

func main() {
	esAddress := "http://localhost:9200"
	esUser := "elastic"
	esPass := "hackathon123"

	esIndex := "avinode_api"

	es := Es(esAddress, esUser, esPass)

	// Example query
	query := "How should I be working with our APIs?"

	// Get documents matching the query
	res, err := es.Get(esIndex, query, 2)
	if err != nil {
		log.Fatalf("Error performing search: %s", err)
	}
	defer res.Body.Close()

	// Print the response
	fmt.Println(res.String())
}
