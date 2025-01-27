package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"linear-cli/api"
)

var message string
var status string

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issue called")

		fmt.Println(message)
		result, err := api.CreateIssue(message, "Backlog")
		if err != nil {
			fmt.Println(err)
		}

		// make this with colors
		fmt.Println("Creating issue...")
		fmt.Println("branch name:", result.BranchName)
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	issueCmd.Flags().StringVarP(&message, "message", "m", "", "Title of the issue")
	issueCmd.Flags().StringVarP(&status, "status", "s", "", "status of the issue (backlog is default)")
}
