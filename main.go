package main

import (
	"linear-cli/cmd"
	_ "linear-cli/cmd/issue"
	_ "linear-cli/cmd/issue/create"
	_ "linear-cli/cmd/issue/list"
)

func main() {
	cmd.Execute()
}
