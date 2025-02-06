package list

import (
	"fmt"
	"linear-cli/api"
	"linear-cli/cmd/issue"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

func headerStyle(n int) lipgloss.Style {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7B8098")).
		Underline(true).UnderlineSpaces(true).
		Width(n)
	return style
}

type Row struct {
	identifier string
	name       string
	status     string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		result, err := api.ListIssues()
		if err != nil {
			fmt.Println(err)
		}

		var rows [][]string

		for i := 0; i < len(result.Team.Issues.Nodes); i++ {
			row := []string{
				result.Team.Issues.Nodes[i].Identifier,
				result.Team.Issues.Nodes[i].Title,
				// lipgloss.JoinVertical(lipgloss.Left, result.Team.Issues.Nodes[i].Title, result.Team.Issues.Nodes[i].BranchName),
				result.Team.Issues.Nodes[i].State.Name,
				result.Team.Issues.Nodes[i].BranchName,
			}
			rows = append(rows, row)
		}

		ta := table.New().
			BorderHeader(false).
			BorderBottom(false).
			BorderTop(false).
			BorderLeft(false).
			BorderRight(false).
			BorderColumn(false).
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#7B8098"))).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch {
				case row == table.HeaderRow:
					if col == 0 {
						return headerStyle(10)
					}
					if col == 1 {

						return headerStyle(68)
					}
					return headerStyle(24)
				case col == 2:
					if rows[row][col] == "Todo" {

						return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#E2E2E2"))
					}

					return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#F0BF00"))
				case col == 0:
					return lipgloss.NewStyle().Foreground(lipgloss.Color("#A5E2A1"))
				default:
					return lipgloss.NewStyle()
				}
			}).
			// Priority, updated_at, branch name
			Headers("ID", "TITLE", "STATUS", "BRANCH").
			Rows(rows...)

		fmt.Println(ta)
	},
}

func init() {
	issue.IssueCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolP("all", "a", true, "List all issues instead of only your issues")
}
