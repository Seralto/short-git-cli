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
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// cCmd represents the c command
var cCmd = &cobra.Command{
	Use:   "c Commit message",
	Short: "A short way of add all files and commit",

	Run: func(cmd *cobra.Command, args []string) {
		message := strings.Join(args, " ")

		exec.Command("git", "add", ".").Run()
		exec.Command("git", "commit", "-m", message).Run()
	},
}

func init() {
	rootCmd.AddCommand(cCmd)
}
