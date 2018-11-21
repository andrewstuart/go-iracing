// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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