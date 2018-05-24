package cmd

import (
	"errors"
	"fmt"
	"os"

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
		db, err := LoadCity()
		if err != nil {
			fmt.Printf("IP datx load error: %s\n", err)
			return
		}
		ipAddr := args[0]
		l, err := db.Find(ipAddr)
		if err != nil {
			fmt.Printf("IP find error: %s\n", err)
			return
		}
		fmt.Printf("%#v", l)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func LoadCity() (*ip.City, error) {
	return ip.NewCity("./data/17monipdb.datx")
}
