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
	"log"
	"os"
	"strings"

	"github.com/scbizu/mew/filter"
	"github.com/scbizu/mew/linker"
	"github.com/spf13/cobra"
)

var repoName string
var gopath string
var grep string
var excludeDirs []string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mew",
	Short: "mew - Show your Go repo related pkgs",
	Long:  `mew - Show your Go repo related pkgs`,
	Run: func(cmd *cobra.Command, args []string) {
		l := linker.NewLinker(gopath, repoName)
		pkgs, err := l.GetAllPkgNames(false, excludeDirs)
		if err != nil {
			log.Fatalln(err)
		}
		pkgFilter := filter.NewFilter(pkgs)
		pkgs = pkgFilter.Grep(grep)
		fmt.Println(pkgs)
		return
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalln("Empty GOPATH,plz set your GOPATH and try again")
	}
	if strings.HasSuffix(repoName, "/") || strings.HasPrefix(repoName, "/") {
		log.Fatalln("Invalid RepoPath")
	}
	RootCmd.Flags().StringVarP(&repoName, "repo", "r", "", "input repo name")
	RootCmd.Flags().StringVarP(&grep, "grep", "g", "", "grep the pkg list")
	RootCmd.Flags().StringArrayVarP(&excludeDirs, "ed", "e", []string{"vendor", ".git"}, "exclude the dir")
}
