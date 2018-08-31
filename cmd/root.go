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
	"os"

	"github.com/spf13/cobra"
	"github.com/ngaut/log"
	"syscall"
	"github.com/daiguadaidai/go-d-bus-password-generator/service"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-d-bus-password-generator",
	Short: "对go-d-bus相关字符加密解密工具",
	Long: `
    对相关字符串进行加密解密,

    # 加密
    ./go-d-bus-password-generator encrypt [需要加密的字符串]
    ./go-d-bus-password-generator encrypt oracle

    # 解密 
    ./go-d-bus-password-generator decrypt [需要解密的字符串]
    ./go-d-bus-password-generator decrypt 6fd887932b8419a1ed24b3055cdef5da42382abcc3a427ed11b31afa36d6bcee
    `,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// 加密数据的子命令
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "加密数据",
	Long: ` 
    对给定的字符串进行加密:

    ./go-d-bus-password-generator encrypt [需要加密的字符串]
    ./go-d-bus-password-generator encrypt oracle
    `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Errorf("没有找到需要加密的字符串. 请确认您输入的需要加密的是否正确性.")
			syscall.Exit(1)
		}

		service.Encrypt(args[0])
	},
}

// 解密数据的子命令
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "解密数据",
	Long: ` 
    对给定的字符串进行解密:

    ./go-d-bus-password-generator decrypt [需要解密的字符串]
    ./go-d-bus-password-generator decrypt 6fd887932b8419a1ed24b3055cdef5da42382abcc3a427ed11b31afa36d6bcee
    `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Errorf("没有找到需要解密的字符串. 请确认您输入的需要加密的是否正确性.")
			syscall.Exit(1)
		}

		service.Decrypt(args[0])
	},
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
	// 添加 encrypt, decrypt 子命令
	rootCmd.AddCommand(encryptCmd, decryptCmd)
}
