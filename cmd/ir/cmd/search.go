package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for an iRacing racer",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		res, err := c.Search(strings.Join(args, " "))
		if err != nil {
			log.Fatal(err)
		}

		// var wg sync.WaitGroup
		// wg.Add(len(res.Racers))
		for i := range res.Racers {
			// go func(i int) {
			// 	defer wg.Done()

			r := res.Racers[i]
			// mem, err := c.GetMember(r.CustID)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			fmt.Printf("%s, logged in %s\n", r.Name /*mem.MemberSince, */, r.LastLogin)
			// }(i)
		}
		// wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
