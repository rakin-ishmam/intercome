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
	"fmt"

	"github.com/rakin-ishmam/intercome/invite"

	"github.com/spf13/cobra"
)

var fileLoc string
var maxDist float64

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list will print customer list",
	Long:  `list will print customer list`,
	Run: func(cmd *cobra.Command, args []string) {
		inv := invite.NewInviter()
		inv.Read(fileLoc)

		if inv.Err() != nil {
			fmt.Println("error", inv.Err())
		}

		flt := invite.Filter{}
		flt.SetOffLoc(53.339428, -6.257664)
		flt.SetMaxDist(maxDist)

		custs := inv.List(flt)

		for _, c := range custs {
			fmt.Printf("name: %s, user id: %d\n", c.Name, c.UserID)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&fileLoc, "file", "f", "./customer.json", "file location")
	listCmd.Flags().Float64VarP(&maxDist, "dist", "d", 100, "maximum distance from office")
}
