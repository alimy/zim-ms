// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alimy/kry/internal/generator"
	"github.com/spf13/cobra"
)

var (
	dstPath  string
	category string
	target   string
	name     string
	pkgName  string
	style    []string
)

func init() {
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "create template project",
		Long:  "create template project",
		Run:   newRun,
	}

	// parse flags for agentCmd
	newCmd.Flags().StringVarP(&dstPath, "dst", "d", ".", "genereted destination target directory")
	newCmd.Flags().StringVarP(&name, "name", "n", "zim", "custom service app name")
	newCmd.Flags().StringVarP(&category, "category", "c", "service", "custom service category")
	newCmd.Flags().StringVarP(&target, "target", "t", "calculator", "custom target service name")
	newCmd.Flags().StringVarP(&pkgName, "pkg", "p", "gitbus.com/exlab/zim-ss/app/service/calculator", "project's package name")
	newCmd.Flags().StringSliceVarP(&style, "style", "s", []string{"tars"}, "generated engine type style:[{tars[,mir]]")

	// register agentCmd as sub-command
	register(newCmd)
}

// newRun run new command
func newRun(_cmd *cobra.Command, _args []string) {
	path, err := filepath.EvalSymlinks(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			if !filepath.IsAbs(dstPath) {
				cwd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				path = filepath.Join(cwd, dstPath)
			} else {
				path = dstPath
			}
		} else {
			log.Fatal(err)
		}
	}
	if err = generator.Generate(path, style, name, category, target, pkgName); err != nil {
		log.Fatal(err)
	}
}
