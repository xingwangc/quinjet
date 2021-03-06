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
	"github.com/spf13/cobra"
	"github.com/xingwangc/quinjet/pkg/config"
	"github.com/xingwangc/quinjet/pkg/rabbitmq"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [task response]",
	Short: "Check the status of CI job",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := rabbitmq.QueryMsgQueue(
			config.GlobalCFG.RabbitMQ.Host,
			config.GlobalCFG.RabbitMQ.Port,
			config.GlobalCFG.RabbitMQ.User,
			config.GlobalCFG.RabbitMQ.Password,
			args[0],
			"unsync",
		)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
