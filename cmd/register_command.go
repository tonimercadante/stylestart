package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var registerCommand = &cobra.Command{
	Use:   "register",
	Short: "Register a command",
	Long:  `Register a command to a workspace`,
	Run:   runRegisterCommand,
}

func runRegisterCommand(cmd *cobra.Command, args []string) {
	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	var projectPath string
	fmt.Printf("Project path: (%s) >", currentDir)
	fmt.Scanln(&projectPath)

	if projectPath == "" {
		projectPath = currentDir
	}

	fmt.Print("Commands to execute(separate by comma to add many): ")

	var commandExec string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		commandExec = scanner.Text()
	}

	fmt.Println(commandExec)

	var commandAlias string
	fmt.Print("Alias for this project: ")
	fmt.Scanln(&commandAlias)

	workspaceJSONname := fmt.Sprintf("data/%s.json", workspaceName)
	workspaceJSON, _ := os.ReadFile(workspaceJSONname)

	var workspace map[string]interface{}

	if err := json.Unmarshal(workspaceJSON, &workspace); err != nil {
		panic(err)
	}

	space := &Spaces{
		Alias:    commandAlias,
		Dir:      projectPath,
		Commands: strings.Split(commandExec, ","),
	}

	spaces := workspace["spaces"].([]interface{})
	spaces = append(spaces, space)
	workspace["spaces"] = spaces

	spaceJson, _ := json.Marshal(workspace)

	fileName := fmt.Sprintf("data/%s.json", workspaceName)

	_ = os.WriteFile(fileName, spaceJson, 0644)
}

func init() {

}
