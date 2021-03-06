// Copyright © 2018 scnace scbizu@gmail.com
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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/scbizu/mew/drawer"
	"github.com/scbizu/mew/filter"
	"github.com/scbizu/mew/linker"
	"github.com/scylladb/go-set/strset"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var repoName string
var gopath string
var grep string
var excludeDirs []string
var isShowJSON bool
var dumpGraph string
var deepMode bool
var dir string
var dirEx []string
var short bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mew",
	Short: "mew - Show your Go repo related pkgs",
	Long:  `mew - Show your Go repo related pkgs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repoNames := []string{repoName}
		if dir != "" {
			repoNames = []string{}
			s := strset.New(dirEx...)
			// ignore the repoName
			if err := filepath.Walk(dir, func(path string, info os.FileInfo, _ error) error {
				if info.IsDir() && !s.Has(info.Name()) {
					repoNames = append(repoNames, filepath.Clean(strings.TrimPrefix(path, gopath+"/src/")))
				}
				return nil
			}); err != nil {
				return err
			}
		}
		d := drawer.NewDot()
		for _, repoName := range repoNames {
			l := linker.NewLinker(gopath, repoName)
			if deepMode {
				if jsonRes := handlePKGMap(l, excludeDirs, grep, dumpGraph, repoName, isShowJSON); jsonRes != "" {
					fmt.Println(jsonRes)
				}
			} else {
				pkgs, err := getpkgs(l, excludeDirs, grep, dumpGraph, repoName)
				if err != nil {
					return err
				}
				repoName = shortFunc()(repoName)
				if err := d.AddDep(repoName, pkgs); err != nil {
					return err
				}
			}
		}
		if err := d.WriteFile(dumpGraph); err != nil {
			return err
		}
		return nil
	},
}

func shortFunc() func(string) string {
	return func(s string) string {
		if short {
			return filepath.Base(s)
		}
		return s
	}
}

func handlePKGMap(l *linker.Linker, excludeDirs []string, grep string, graphName string, repo string, isShowJSON bool) string {
	var pkgMap map[string][]string
	var err error
	pkgMap, err = l.GetAllPKGNames(false, excludeDirs)
	if err != nil {
		logrus.Fatalln(err)
	}

	pkgMapFilter := filter.NewMapFilter(pkgMap)
	pkgMapFilter.Grep(grep)

	if err = drawer.DrawWithMapAndSave(repo, graphName, pkgMap); err != nil {
		logrus.Fatalln(err.Error())
	}

	if isShowJSON {
		jsonRes, err := json.Marshal(pkgMap)
		if err != nil {
			logrus.Fatalln(err)
		}
		return string(jsonRes)
	}
	return ""
}

func getpkgs(l *linker.Linker, excludeDirs []string, grep string, graphName string, repo string) ([]string, error) {
	var pkgs []string
	var err error
	pkgs, err = l.GetLayerPKGNames(false, excludeDirs)
	if err != nil {
		return nil, err
	}
	pkgFilter := filter.NewFilter(pkgs)
	pkgs = pkgFilter.Grep(grep)
	for index := range pkgs {
		pkgs[index] = shortFunc()(pkgs[index])
	}
	return pkgs, nil
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
	RootCmd.Flags().StringVar(&dir, "dir", os.Getenv("GOPATH"), "the whole directory")
	RootCmd.Flags().StringArrayVarP(&excludeDirs, "ed", "e", []string{"vendor", ".git"}, "exclude the dir")
	RootCmd.Flags().StringArrayVar(&dirEx, "ex", []string{"vendor"}, "exclude the dir")
	RootCmd.Flags().BoolVar(&isShowJSON, "json", false, "show json format")
	RootCmd.Flags().BoolVar(&short, "short", false, "file base name")
	RootCmd.Flags().BoolVar(&deepMode, "deep", false, "[Experimental feature]in deep mode,you will get all(include really all dependency) third party related pkg name")
	RootCmd.Flags().StringVarP(&dumpGraph, "graph", "d", drawer.DefaultFileName, "dump graphviz graph")
}
