package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"os"
)

const baseURL = "https://api.linear.app/graphql"

var client *resty.Client
var teamId string

func init() {
	client = resty.New()
	fmt.Println("asdasdasd")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiKey := os.Getenv("LINEAR_API_KEY")
	if apiKey == "" {
		panic("LINEAR_API_KEY environment variable is not set")
	}

	teamId = os.Getenv("TEAM_ID")
	if teamId == "" {
		panic("TEAM_ID environment variable is not set")
	}

	client.SetHeader("Authorization", apiKey)
	client.SetHeader("Content-Type", "application/json")
}

// GraphQLRequest represents a GraphQL query or mutation request
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLResponse represents a GraphQL response
type GraphQLResponse struct {
	Data   interface{}    `json:"data"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

// GraphQLError represents a GraphQL error
type GraphQLError struct {
	Message string   `json:"message"`
	Path    []string `json:"path,omitempty"`
}

// ExecuteMutation performs a GraphQL mutation and unmarshals the response into the provided result interface
func ExecuteMutation(mutation string, variables map[string]interface{}, result interface{}) error {
	// Create the request body
	requestBody := GraphQLRequest{
		Query:     mutation,
		Variables: variables,
	}

	// Make the request
	resp, err := client.R().
		SetBody(requestBody).
		Post(baseURL)

	if err != nil {
		return fmt.Errorf("failed to execute mutation: %w", err)
	}

	// Parse the response
	var graphQLResp GraphQLResponse
	if err := json.Unmarshal(resp.Body(), &graphQLResp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for GraphQL errors
	if len(graphQLResp.Errors) > 0 {
		return fmt.Errorf("graphql errors: %v", graphQLResp.Errors)
	}

	// Unmarshal the data into the result
	dataBytes, err := json.Marshal(graphQLResp.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, result); err != nil {
		return fmt.Errorf("failed to unmarshal response data into result: %w", err)
	}

	return nil
}
