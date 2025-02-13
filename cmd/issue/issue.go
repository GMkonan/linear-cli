package issue

import (
	"fmt"
	"linear-cli/cmd"

	"github.com/spf13/cobra"
)

// issueCmd represents the issue command
var IssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "A brief description of your commandsdasdasdasd",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issue called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(IssueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
