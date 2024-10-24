package cmd

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/spf13/cobra"
)

var startWorkspace = &cobra.Command{
	Use:   "start",
	Short: "Start a workspace",
	Long:  `Execute all the commands needed to start your workspace environment, including deps(init a database)`,
	Run:   runStartWorkspace,
}

func runStartWorkspace(cmd *cobra.Command, args []string) {
	log.Println("Starting your work environment")
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// command := "echo 'Hello from new terminal!'"
	// terminal := exec.Command("osascript", "-e", `tell application "Terminal" to do script "`+command+`"`)

	// terminal.Start()
}
