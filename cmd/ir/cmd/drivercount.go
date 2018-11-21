package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// drivercountCmd represents the drivercount command
var drivercountCmd = &cobra.Command{
	Use:   "drivercount",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		ct, err := c.GetDriverCounts()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Total: %d, Mine: %d, Total Laps: %s\n", ct.Total, ct.Myracers, ct.LapCount)
	},
}

func init() {
	rootCmd.AddCommand(drivercountCmd)
}
