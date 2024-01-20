/*
Copyright © 2024 Ufuk Bombar ufukbombar@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gofee",
	Short: "gofee (Golang Financial Event Engine) is an event engine to process financial events.",
	Long: `gofee (Golang Financial Event Engine) is an event engine to process 
financial events. To create a database please sue the init command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gofee, gofee help to display help.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
