package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

const baseURL = "https://api.linear.app/graphql"

func main() {
	api_key := os.Getenv("LINEAR_API_KEY")
	client := resty.New()
	client.SetHeader("Authorization", api_key)
	client.SetHeader("Content-Type", "application/json")
}

func createIssue(client *resty.Client, title string, description string) {
	mutation := fmt.Sprintf(`
    mutation {
        issueCreate(input: {title: "%s", description: "%s"}) {
            success
            issue {
                id
                title
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

	fmt.Println(resp.String())
}
