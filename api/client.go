package api

import (
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
