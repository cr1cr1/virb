/*

  Copyright 2019 Google LLC

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package cmd

import (
	"github.com/scottlaird/virb"
	"github.com/spf13/cobra"
)

// deletefilegroupCmd represents the deletefilegroup command
var deletefilegroupCmd = &cobra.Command{
	Use:   "deletefilegroup",
	Short: "Delete a video file and all of its associated metadata",
	Long:  `deletefilegroup <filename>...`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := virb.DeleteFileGroup(hostname, args)
		if err != nil {
			panic(err)
		}
		printResponse(resp)
	},
}

func init() {
	rootCmd.AddCommand(deletefilegroupCmd)
}
