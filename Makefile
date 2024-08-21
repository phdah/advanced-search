
run: build
	bin/advanced-search "What does an aircraft refer to, and what types exists?"

build:
	go build -o ./bin/advanced-search ./cmd/advanced-search.go

elastic-search:
	ELASTICSEARCH_PASSWORD="hackathon123" && $$HOME/repos/privat/insta-infra/run.sh elasticsearch
	curl -u "elastic:hackathon123" -X GET "http://localhost:9200"

scrape: build-scrape
	go run ./cmd/scraping-docs.go

build-scrape:
	go build -o ./bin/scraping-docs ./cmd/scraping-docs.go
