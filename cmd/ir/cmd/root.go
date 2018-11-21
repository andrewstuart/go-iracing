package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ir",
	Short: "Interact with the iRacing API from your command line",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ir.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var home string

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		var err error
		// Find home directory.
		home, err = homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ir" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(path.Join(home, ".config"))
		viper.SetConfigName(".iracing")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	viper.ReadInConfig()

	viper.AutomaticEnv() // read in environment variables that match
}
