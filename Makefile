
run: build
	./bin/advanced-search

build:
	go build -o ./bin/advanced-search ./cmd/advanced-search.go

elastic-search:
	ELASTICSEARCH_PASSWORD="hackathon123" && $$HOME/repos/privat/insta-infra/run.sh elasticsearch
	curl -u "elastic:hackathon123" -X GET "http://localhost:9200"

scrape:
	go run ./cmd/scraping-docs.go
