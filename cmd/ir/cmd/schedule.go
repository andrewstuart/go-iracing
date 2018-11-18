package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"astuart.co/iracing"
	"github.com/spf13/cobra"
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Get the iRacing schedule",
	Run: func(cmd *cobra.Command, args []string) {
		c := iracing.Client{}

		s, err := c.GetSchedule()
		if err != nil {
			log.Fatal(err)
		}

		sort.Slice(s.Contents, func(i, j int) bool {
			return s.Contents[i].EventAt.Before(s.Contents[j].EventAt.Time)
		})

		tw := tabwriter.NewWriter(os.Stdout, 0, 3, 1, ' ', 0)
		_, err = fmt.Fprintln(tw, "Event\tStart")
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range s.Contents {
			if c.EventAt.Before(time.Now()) || c.EventAt.After(time.Now().Add(12*time.Hour)) {
				continue
			}
			_, err = fmt.Fprintf(tw, "%s\t%s\n", plusReplacer.Replace(c.Bannertext), c.EventAt.Format(time.RFC1123))
			if err != nil {
				log.Fatal(err)
			}
		}
		err = tw.Flush()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
