package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"os"
)

const baseURL = "https://api.linear.app/graphql"

var client *resty.Client

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

	client.SetHeader("Authorization", apiKey)
	client.SetHeader("Content-Type", "application/json")
}

func CreateIssue(title string, description string) string {
	mutation := fmt.Sprintf(`
    mutation {
        issueCreate(input: {title: "%s", description: "%s"}) {
            success
            issue {
                id
                title
		branchName
            }
        }
    }
    `, title, description)

	resp, err := client.R().
		SetBody(map[string]string{"query": mutation}).
		Post(baseURL)

	if err != nil {
		panic(err)
	}

	// fmt.Println(resp.String())
	return resp.String()
}
