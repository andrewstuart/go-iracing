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
	"github.com/spf13/viper"
)

// racesCmd represents the races command
var racesCmd = &cobra.Command{
	Use:   "races",
	Short: "Get a list of races today",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := iracing.NewClient()
		if err != nil {
			log.Fatal(err)
		}
		err = c.Login(viper.GetString("iracing.user"), viper.GetString("iracing.password"))
		if err != nil {
			log.Fatal(err)
		}

		guide, err := c.GetRaceGuide()
		if err != nil {
			log.Fatal(err)
		}

		sort.Slice(guide.Series, func(i, j int) bool {
			ni, nj := guide.Series[i].CurrentSchedule(), guide.Series[j].CurrentSchedule()
			if ni == nil {
				return false
			}
			if nj == nil {
				return true
			}

			ri, rj := ni.NextRace(), nj.NextRace()
			if ri == nil {
				return false
			}
			if rj == nil {
				return true
			}

			return ri.StartTime.Before(rj.StartTime.Time)
		})

		tw := tabwriter.NewWriter(os.Stdout, 0, 3, 2, ' ', 0)
		fmt.Fprintln(tw, "Series\tTrack\tStart\tUntil\tLength")

		for _, series := range guide.Series {
			if !series.Eligible {
				continue
			}
			ns := series.CurrentSchedule()
			if ns == nil {
				break
			}
			nr := ns.NextRace()
			if nr == nil {
				break
			}
			dur := hmString(nr.EndTime.Sub(nr.StartTime.Time))
			// if nr.RaceTimeLimitMinutes <= 0 {
			// 	dur = fmt.Sprintf("%d laps", nr.RaceLapLimit)
			// }
			fmt.Fprintf(
				tw,
				"%s\t%s\t%s\t%s\t%s\n",
				series.SeriesName,
				nr.TrackName,
				nr.StartTime.Format(time.RFC1123),
				hmString(time.Until(nr.StartTime.Time)),
				dur,
			)
		}
		tw.Flush()
	},
}

func hmString(d time.Duration) string {
	dh := d.Truncate(time.Hour)
	h := dh.Hours()
	m := (d - dh).Truncate(time.Minute).Minutes()

	return fmt.Sprintf("%02.0f:%02.0f", h, m)
}

func init() {
	rootCmd.AddCommand(racesCmd)

	racesCmd.Flags().StringP("username", "u", "", "iRacing username")
	viper.BindPFlag("iracing.username", racesCmd.Flags().Lookup("username"))

	racesCmd.Flags().StringP("password", "p", "", "iRacing password")
	viper.BindPFlag("iracing.password", racesCmd.Flags().Lookup("password"))
}
