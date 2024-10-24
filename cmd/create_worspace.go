package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Spaces struct {
	Name     string   `json:"name"`
	Dir      string   `json:"dir"`
	NewTab   bool     `json:"newtab"`
	Commands []string `json:"commands"`
}

type space struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Spaces      []Spaces `json:"spaces"`
}

var name string

var createWorkspace = &cobra.Command{
	Use:   "create",
	Short: "Create a workspace",
	Long:  `Execute all the commands needed to start your workspace environment, including deps(init a database)`,
	Run:   runCreateWorkspace,
}

func runCreateWorkspace(cmd *cobra.Command, args []string) {
	log.Println("Creating new workspace", args, name)

	space := &space{
		Name:   name,
		Spaces: []Spaces{},
	}

	spaceJson, _ := json.Marshal(space)

	fileName := fmt.Sprintf("data/%s.json", name)
	_ = os.WriteFile(fileName, spaceJson, 0644)

	fmt.Println(string(spaceJson))
}

func init() {
	createWorkspace.Flags().StringVarP(&name, "name", "n", "", "Name of the workspace")
	createWorkspace.MarkFlagRequired("name")
}
