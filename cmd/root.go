package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stylestart",
	Short: "Stylestart is a simple Cobra application",
	Long:  `A simple command-line application using Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome to Stylestart. Start your workspace with style!")
		log.Println("Hugs are worth more than handshakes!")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startWorkspace)
}
