package cmd

import (
	"log"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var memberCmd = &cobra.Command{
	Use:   "member",
	Args:  cobra.ExactArgs(1),
	Short: "Get details about a specified iRacing member",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		member, err := c.GetMember(i)
		if err != nil {
			log.Fatal(err)
		}

		spew.Dump(member)
	},
}

func init() {
	rootCmd.AddCommand(memberCmd)

}
