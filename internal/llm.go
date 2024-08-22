package internal

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
)

type LLMContext struct {
    Prompt      string
    Response    string
    Context     string
}

type LLMResponse struct {
    // {"model":"llama3.1","created_at":"2024-08-21T13:17:28.070757Z","response":" of","done":false}
    // Model     string `json:"model"`
    // CreatedAt string `json:"created_at"`
    Response string `json:"response"`
    // Done      string `json:"done"`
}

func AskOllamaQuestion(llmContext LLMContext) (LLMContext, error) {
    url := "http://localhost:11434/api/generate"
    model := "llama3.1"

    // Prepare the request body as a map
    llmContext.Context += "\nQuestion: " + llmContext.Prompt
    requestBody := map[string]string{
        "model":  model,
        "prompt": llmContext.Context,
    }

    // Convert the request body to JSON
    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return llmContext, fmt.Errorf("error marshalling request body: %v", err)
    }

    // Create a new HTTP request
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return llmContext, fmt.Errorf("error creating request: %v", err)
    }

    // Set the appropriate headers
    req.Header.Set("Content-Type", "application/json")

    // Send the request using the http client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return llmContext, fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return llmContext, fmt.Errorf("error reading response body: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        return llmContext, fmt.Errorf("received non-OK response: %s", body)
    }

    // Convert the body to a string and split by newlines to handle multiple JSON objects
    rawResponses := strings.Split(string(body), "\n")

    // Initialize a variable to store the combined response
    var combinedResponse string

    // Iterate over each raw JSON object
    for _, raw := range rawResponses {
        if strings.TrimSpace(raw) == "" {
            continue // Skip empty lines
        }

        var responseObj LLMResponse
        if err := json.Unmarshal([]byte(raw), &responseObj); err != nil {
            return llmContext, fmt.Errorf("error unmarshalling response: %v", err)
        }

        // Append each part of the response to the combined response
        combinedResponse += responseObj.Response
    }

    llmContext.Context += "\nResponse: " + combinedResponse
    llmContext.Response = combinedResponse
    // Return the combined response
    return llmContext, nil
}
