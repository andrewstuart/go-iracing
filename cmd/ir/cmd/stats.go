package cmd

import (
	"log"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Stats for a driver",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		c, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		stats, err := c.GetStats(i)
		if err != nil {
			log.Fatal(err)
		}

		spew.Dump(stats)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
