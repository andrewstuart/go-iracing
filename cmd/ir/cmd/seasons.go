package cmd

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// seasonsCmd represents the seasons command
var seasonsCmd = &cobra.Command{
	Use:   "seasons",
	Short: "print out a list of seasons",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		seasons, err := c.GetSeasons(nil)
		if err != nil {
			log.Fatal(err)
		}

		spew.Dump(seasons)
	},
}

func init() {
	rootCmd.AddCommand(seasonsCmd)
}
