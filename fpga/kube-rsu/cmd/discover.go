// Copyright 2019 Intel Corporation. All rights reserved
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

package rsu

import (
	"fmt"
	"github.com/spf13/cobra"
)

// discoverCmd represents the discover command
var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover FPGA",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("discover called!")
	},
}

func init() {

	const help = `Discover FPGA

Usage:
  rsu discover

Flags:
  -h, --help       help
`
	// add `discover` command
	rsuCmd.AddCommand(discoverCmd)
	discoverCmd.SetHelpTemplate(help)
}
