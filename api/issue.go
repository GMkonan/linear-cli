package api

import (
	// "encoding/json"
	"fmt"
	"log"
	"time"
)

type Issue struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	BranchName string `json:"branchName"`
}

type IssueCreate struct {
	IssueCreate struct { // Nested struct matching the operation name
		Success bool  `json:"success"`
		Issue   Issue `json:"issue"`
	} `json:"issueCreate"` // Field name matches the GraphQL operation
}

func CreateIssue(title string) (*Issue, error) {
	mutation := `
	mutation IssueCreate($teamId: String!, $title: String!) {
 		issueCreate(input: {teamId: $teamId ,title: $title}) {
             success
             issue {
                 id
                 title
 		branchName
             }
         }
     }
    `

	variables := map[string]interface{}{
		"teamId": teamId,
		"title":  title,
	}

	var result IssueCreate
	err := GraphQL(mutation, variables, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	fmt.Printf("Created issue with ID: %s\n", result.IssueCreate.Issue.ID)
	return &result.IssueCreate.Issue, nil
}

type Team struct {
	Team struct { // Nested struct matching the operation name
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
	} `json:"Team"` // Field name matches the GraphQL operation
}

func ListIssues() (*Team, error) {

	query := `
	query Team($id: String!) {
  team(id: $id) {
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
}`
	variables := map[string]interface{}{
		"id": teamId,
	}

	var result Team
	err := GraphQL(query, variables, &result)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(result)

	return &result, nil
}
