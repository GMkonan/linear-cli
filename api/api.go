package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const baseURL = "https://api.linear.app/graphql"

var client *resty.Client
var teamId string

func init() {
	client = resty.New()

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

type Issue struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	BranchName string `json:"branchName"`
}

func CreateIssue(title string) (*Issue, error) {

	type GraphqlResponse struct {
		Data struct {
			IssueCreate struct {
				Success bool `json:"success"`
				Issue   struct {
					ID         string `json:"id"`
					Title      string `json:"title"`
					BranchName string `json:"branchName"`
				} `json:"issue"`
			} `json:"issueCreate"`
		} `json:"data"`
	}

	mutation := fmt.Sprintf(`
    mutation {
        issueCreate(input: {teamId: "%s" ,title: "%s"}) {
            success
            issue {
                id
                title
		branchName
            }
        }
    }
    `, teamId, title)

	resp, err := client.R().
		SetBody(map[string]string{"query": mutation}).
		Post(baseURL)

	if err != nil {
		panic(err)
	}

	var response GraphqlResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return &Issue{}, fmt.Errorf("Failed to parse response: %w", err)
	}

	if !response.Data.IssueCreate.Success {
		return &Issue{}, fmt.Errorf("Issue creation failed")
	}

	issue := response.Data.IssueCreate.Issue

	return &Issue{
		ID:         issue.ID,
		Title:      issue.Title,
		BranchName: issue.BranchName,
	}, nil
}
