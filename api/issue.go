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
				ID         string `json:"id"`
				Identifier string `json:"identifier"`
				Title      string `json:"title"`
				BranchName string `json:"branchName"`
				State      struct {
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
	query Team($id: String!, $userId: ID, $status: [String!]) {
  team(id: $id) {
    id
    name
    issues(filter:  {
       assignee:  {
          id:  {
             eq: $userId
          }
       },
       state:  {
        or: [ {
          name: {
            in: $status
          }
          }]
       }
    }) {
      nodes {
        id
	identifier
        title
	branchName
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

	myIssuesOnly := true

	status := []string{"Todo", "In Progress"}

	variables := map[string]interface{}{
		"id":     teamId,
		"status": status,
	}

	if myIssuesOnly == true {
		// variables["userId"] = "08c25e3a-0e0a-413a-bf5d-78f0e2ed4618"

		// Work
		variables["userId"] = "946fed3f-7d25-4279-be1f-068198649b94"
	}

	var result Team
	err := GraphQL(query, variables, &result)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(result)

	return &result, nil
}
