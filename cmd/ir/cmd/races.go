package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"astuart.co/iracing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// racesCmd represents the races command
var racesCmd = &cobra.Command{
	Use:   "races",
	Short: "Get a list of races today",
	Run: func(cmd *cobra.Command, args []string) {
		c := iracing.Client{}
		err := c.Login(viper.GetString("iracing.user"), viper.GetString("iracing.password"))
		if err != nil {
			log.Fatal(err)
		}

		guide, err := c.GetRaceGuide()
		if err != nil {
			log.Fatal(err)
		}

		tw := tabwriter.NewWriter(os.Stdout, 0, 3, 1, ' ', 0)
		fmt.Fprintln(tw, "Series\tTrack\tStart")

		for _, series := range guide.Series {
			if !series.Eligible {
				continue
			}

			var latest *iracing.SeasonSchedule
			for _, sched := range series.SeasonSchedules {
				if latest == nil || sched.SeasonStartDate.After(latest.SeasonStartDate.Time) {
					latest = &sched
					continue
				}
			}

			if len(latest.Races) == 0 {
				continue
			}

			sort.Slice(latest.Races, func(i, j int) bool {
				return latest.Races[i].StartTime.Before(latest.Races[j].StartTime.Time)
			})

			lr := latest.Races[len(latest.Races)-1]
			fmt.Fprintf(tw, "%s\t%s\t%s\n", dePlus(series.SeriesName), dePlus(lr.TrackName), lr.StartTime)
		}
		tw.Flush()
	},
}

func init() {
	rootCmd.AddCommand(racesCmd)

	racesCmd.Flags().StringP("username", "u", "", "iRacing username")
	viper.BindPFlag("iracing.username", racesCmd.Flags().Lookup("username"))

	racesCmd.Flags().StringP("password", "p", "", "iRacing password")
	viper.BindPFlag("iracing.password", racesCmd.Flags().Lookup("password"))
}
