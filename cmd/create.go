/*
Copyright (c) 2014 and onwards Nimble.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// createCmd generates the project using cookiecutter
var (
	branchFlag string

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create the gin project with your own app name.",
		Run: func(cmd *cobra.Command, args []string) {
			shCmd := exec.Command("cookiecutter", "https://github.com/nimblehq/gin-templates.git", "--checkout", branchFlag)
			shCmd.Stdout = os.Stdout
			shCmd.Stdin = os.Stdin

			err := shCmd.Run()
			if err != nil {
				log.Fatal("Gin project created unsuccessfully: ", err)
			} else {
				log.Print("Gin project created successfully.")
			}
		},
	}
)

func init() {
	createCmd.PersistentFlags().StringVarP(&branchFlag, "branch", "b", "main", "template branch")
	rootCmd.AddCommand(createCmd)
}
