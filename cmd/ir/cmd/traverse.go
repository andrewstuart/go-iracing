package cmd

import (
	"fmt"
	"log"
	"strconv"

	iracing "astuart.co/go-iracing"
	"github.com/spf13/cobra"
)

// traverseCmd represents the traverse command
var traverseCmd = &cobra.Command{
	Use:   "traverse",
	Args:  cobra.ExactArgs(1),
	Short: "A demo function to traverse races and winners to a certain depth",
	Run: func(cmd *cobra.Command, args []string) {
		seen := map[int]struct{}{}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		seen[i] = struct{}{}

		c, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		member(c, i, seen, 10)

		fmt.Printf("len(seen) = %+v\n", len(seen))
	},
}

func member(client *iracing.Client, id int, seen map[int]struct{}, depth int) {
	if depth < 0 {
		return
	}
	races, err := client.GetLastRaces(id)
	if err != nil {
		log.Fatal("last races err ", err)
	}

	for _, race := range races {
		if _, ok := seen[race.WinnerID]; !ok {
			fmt.Println(race.WinnerName)
			seen[race.WinnerID] = struct{}{}
			member(client, race.WinnerID, seen, depth-1)
		}
	}
}

func init() {
	rootCmd.AddCommand(traverseCmd)
}
