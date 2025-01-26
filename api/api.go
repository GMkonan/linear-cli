package api

import (
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

func CreateIssue(title string) string {
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

	// fmt.Println(resp.String())
	return resp.String()
}
