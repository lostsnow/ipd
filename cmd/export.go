package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "ipd check",
	Long:  `ipd check`,
	Run: func(cmd *cobra.Command, args []string) {
		city, err := LoadCity()
		if err != nil {
			fmt.Errorf("IP datx load error: %s\n", err)
		}

		fmt.Println("Start check...")

		i, err := city.Check()
		if err != nil {
			fmt.Errorf("IP location check: %s\n", err)
		}

		if len(i.Country) > 0 {
			fmt.Println("Irregular Country:")
			for _, country := range i.Country {
				fmt.Println("  *", country)
			}
		}

		if len(i.State) > 0 {
			fmt.Println("Irregular State:")
			for _, state := range i.State {
				fmt.Println("  *", state)
			}
		}

		if len(i.City) > 0 {
			fmt.Println("Irregular City:")
			for _, city := range i.City {
				fmt.Println("  *", city)
			}
		}

		fmt.Println("End check")
	},
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
