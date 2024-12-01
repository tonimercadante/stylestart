package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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

	fileName := fmt.Sprintf("data/%s.json", workspaceName)
	workspaceJSON, _ := os.ReadFile(fileName)

	var workspace map[string]interface{}

	if err := json.Unmarshal(workspaceJSON, &workspace); err != nil {
		panic(err)
	}

	for _, space := range workspace["spaces"].([]interface{}) {
		spaceMap := space.(map[string]interface{})

		dir := fmt.Sprintf("%v", spaceMap["dir"])
		commands := spaceMap["commands"].([]interface{})

		fullScript := fmt.Sprintf("cd %s", dir)
		for _, cmd := range commands {
			fullScript += fmt.Sprintf("; %v", cmd)
		}

		terminal := exec.Command("osascript", "-e", `tell application "Terminal" to do script "`+fullScript+`"`)
		if err := terminal.Start(); err != nil {
			log.Printf("Error executing commands for %s: %v", spaceMap["alias"], err)
		}
	}

}

func init() {

}
