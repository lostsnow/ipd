package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "ipd api",
	Long:  `ipd api`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api")
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)
}
