package api

import (
	"encoding/json"
	"fmt"
	"time"
)

type Issue struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	BranchName string `json:"branchName"`
}

type IssueCreate struct {
	Success bool  `json:"success"`
	Issue   Issue `json:"issue"`
}

func CreateIssue(title string, state string) (*Issue, error) {

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

	var response GraphqlResponse[IssueCreate]
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return &Issue{}, fmt.Errorf("Failed to parse response: %w", err)
	}

	if !response.Data.Operation.Success {
		return &Issue{}, fmt.Errorf("Issue creation failed")
	}

	issue := response.Data.Operation.Issue

	return &Issue{
		ID:         issue.ID,
		Title:      issue.Title,
		BranchName: issue.BranchName,
	}, nil
}

type Team struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Issues struct {
		Nodes []struct {
			ID    string `json:"id"`
			Title string `json:"title"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
			Assignee      any       `json:"assignee"`
			PriorityLabel string    `json:"priorityLabel"`
			UpdatedAt     time.Time `json:"updatedAt"`
		} `json:"nodes"`
	} `json:"issues"`
}

func ListIssues() (*Team, error) {

	query := fmt.Sprintf(`
    query Team {
  team(id: "%s") {
    id
    name
    issues {
      nodes {
        id
        title
        state {
          id
          name
        }
        assignee {
          id
          name
        }
        priorityLabel
        updatedAt
      }
    }
  }
}`, teamId)
	fmt.Println(query)
	resp, err := client.R().
		SetBody(map[string]interface{}{"query": query}).
		Get(baseURL)
	fmt.Println(resp)
	if err != nil {
		panic(err)
	}

	var response GraphqlResponse[Team]
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return &Team{}, fmt.Errorf("Failed to parse response: %w", err)
	}

	// if !response.Data.Operation.Success {
	// 	return &Team{}, fmt.Errorf("Issue creation failed")
	// }
	fmt.Println(response)
	issues := response.Data.Operation

	return &issues, nil
}
