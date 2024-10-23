package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var startWorkspace = &cobra.Command{
	Use:   "start",
	Short: "Start a workspace",
	Long:  `Execute all the commands needed to start your workspace environment, including deps(init a database)`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Opening...")

	command := "echo 'Hello from new terminal!'"
	terminal := exec.Command("osascript", "-e", `tell application "Terminal" to do script "`+command+`"`)

	terminal.Start()
}
