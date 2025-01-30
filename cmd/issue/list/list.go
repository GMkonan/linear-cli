package list

import (
	"fmt"
	"linear-cli/api"
	"linear-cli/cmd/issue"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func headerStyle(n int) lipgloss.Style {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#808080")).
		Underline(true).UnderlineSpaces(true).
		Width(n)
	return style
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
		fmt.Println("list called")

		result, err := api.ListIssues()
		if err != nil {
			fmt.Println(err)
		}

		// var title = "\033[4mTITLE\033[0m"
		var style = lipgloss.NewStyle().MarginRight(2)

		var title = "TITLE"
		var id = "ID"
		var status = "STATUS"
		var titleZ = result.Team.Issues.Nodes[0].Title
		var idZ = result.Team.Issues.Nodes[0].Identifier
		var statusZ = result.Team.Issues.Nodes[0].State.Name

		var titleStyle = headerStyle(56)
		var idStyle = headerStyle(5)
		var statusStyle = headerStyle(12)

		var i = lipgloss.JoinVertical(lipgloss.Left, idStyle.Render(id), idZ)
		var t = lipgloss.JoinVertical(lipgloss.Left, titleStyle.Render(title), titleZ)
		var s = lipgloss.JoinVertical(lipgloss.Left, statusStyle.Render(status), statusZ)
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left, style.Render(i), t, s))

		for i := 1; i < len(result.Team.Issues.Nodes); i++ {
			fmt.Printf("%s  %s  %s \n", result.Team.Issues.Nodes[i].Identifier, result.Team.Issues.Nodes[i].Title, result.Team.Issues.Nodes[i].State.Name)
		}

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
}
