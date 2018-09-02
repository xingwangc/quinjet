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
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	//	"times"

	"github.com/spf13/cobra"
	"github.com/xingwangc/quinjet/pkg/config"
	"github.com/xingwangc/quinjet/pkg/k8s"
	"github.com/xingwangc/quinjet/pkg/rabbitmq"
)

func citrigger(cmd *cobra.Command, args []string) {
	mode := "unsync"
	if len(args) == 4 && (args[3] == "sync" || args[3] == "unsync") {
		mode = args[3]
	}

	buf := md5.New()
	io.WriteString(buf, strings.Join(args[0:3], " "))
	md5 := buf.Sum(nil)

	err := k8s.CreateCIJob(args[0], args[1], args[2], config.GlobalCFG.Context)
	if err != nil {
		panic(err)
	}

	if mode == "unsync" {
		fmt.Printf("%x\n", md5)
	} else {
		err := rabbitmq.QueryMsgQueue(
			config.GlobalCFG.RabbitMQ.Host,
			config.GlobalCFG.RabbitMQ.Port,
			config.GlobalCFG.RabbitMQ.User,
			config.GlobalCFG.RabbitMQ.Password,
			string(md5),
			mode,
		)
		if err != nil {
			panic(err)
		}
	}

}

// ciCmd represents the ci command
var ciCmd = &cobra.Command{
	Use:   "ci [git repo] [dependance install] [build command] sync/unsync",
	Short: "trigger a ci job",
	Long:  ``,
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ci called")
	},
}

func init() {
	rootCmd.AddCommand(ciCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ciCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ciCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
