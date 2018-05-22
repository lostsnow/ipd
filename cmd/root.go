package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ipipdotnet/datx-go"
	"github.com/lostsnow/ipd/ip"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ipd 1.2.3.4",
	Short: "ipd",
	Long:  `ipd`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires IP address")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		city, err := datx.NewCity("./data/17monipdb.datx")
		if err != nil {
			fmt.Errorf("IP datx load error: ", err)
		}

		ipAddr := args[0]
		ip.Find(city, ipAddr)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
