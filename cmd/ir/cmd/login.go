package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		err = c.Login(viper.GetString("iracing.user"), viper.GetString("iracing.password"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("logged in successfully")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("user", "u", "", "iRacing user name/email")
	viper.BindPFlag("iracing.user", loginCmd.Flags().Lookup("user"))

	loginCmd.Flags().StringP("password", "p", "", "iRacing password")
	viper.BindPFlag("iracing.password", loginCmd.Flags().Lookup("password"))
}
