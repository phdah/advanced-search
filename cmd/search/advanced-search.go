package main

import (
    "fmt"
    "log"
    "os"

    . "github.com/phdah/advanced-search/internal"
)

func main() {
    // Step 1 - Elastic Search
    esAddress := "http://localhost:9200"
    esUser := "elastic"
    esPass := "hackathon123"

    esIndex := "avinode_api"

    es := Es(esAddress, esUser, esPass)

    // Example query:
    // "How should I be working with our APIs?"
    if len(os.Args) <= 1 {
        log.Fatal("Not enough arguments passed. A question is requred")
    }
    query := os.Args[1]

    // Get documents matching the query
    res, err := es.Get(esIndex, query, 2)
    if err != nil {
        log.Fatalf("Error performing search: %s", err)
    }
    defer res.Body.Close()

    document, err := es.Parse(res, "content")
    if err != nil {
        log.Fatalf("Error parsing body: %s", err)
    }
    // Step 2 - Setup prompt
    prompt := `You are a helper with answering questions about documentation. You will be passed; documentation and a question about that. Only answer the given question, using the information in the in the documentation. Never make up an answer.

Format your response like this:

<Document title from where the information is taken>: <the answer to the question(s) preferably using quotes and listing of result>`

    prompt += "\nThis is the documentations: " + document

    prompt += "\nThis is the question: " + query
    // Step 3 - Ask LLM question with documentation
    llmContext := LLMContext{}
    llmContext.Prompt = prompt
    llmContext, err = AskOllamaQuestion(llmContext)
    if err != nil {
        log.Fatalf("Error asking Ollama question: %v", err)
    }

    // Print the llmContext (response)
    fmt.Println(Green + "Model response:\n", Reset + llmContext.Response)

    // Finished
}
