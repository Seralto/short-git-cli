// Copyright © 2019 Sérgio Toledo <contato@sergiotoledo.com.br>
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
	"os/exec"

	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "A short way to push branches",

	Run: func(cmd *cobra.Command, args []string) {
		branch, _ := cmd.Flags().GetString("branch")

		if branch == "m" {
			branch = "master"
		}

		out, err := exec.Command("git", "push", "origin", branch).Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", out)
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.Flags().StringP("branch", "b", "develop", "Branch name")
}
