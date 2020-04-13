// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alimy/mgo/naming"
)

// Generate generate template project
func Generate(dstPath string, style []string, category string, pkgName string, srvName string) error {
	var (
		err               error
		filePath, dirPath string
		file              *os.File
	)

	tmpls, exist := tmplInfosFrom(style)
	if !exist {
		return fmt.Errorf("not exist style(%s) template project", style)
	}

	names := strings.Split(srvName, ".")
	if len(names) != 3 {
		return fmt.Errorf("not availible full service name: %s", srvName)
	}

	slackNaming := naming.NewSnakeNamingStrategy()
	fixedName := slackNaming.Naming(names[1])
	ctx := &tmplCtx{
		AppName:    names[0],
		Category:   category,
		Target:     names[1],
		ObjName:    names[2],
		PkgName:    pkgName,
		SrvName:    srvName,
		PacketName: strings.ReplaceAll(fixedName, "_", "-"),
	}

	tmpl := template.New("kry")
	for fileName, assetInfo := range tmpls {
		filePath = filepath.Join(dstPath, fileName)
		dirPath = filepath.Dir(filePath)
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}

		perm := os.FileMode(0644)
		if assetInfo.isExec {
			perm = 0755
		}
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			break
		}

		if assetInfo.isTmpl {
			t, err := tmpl.Parse(MustAssetString(assetInfo.name))
			if err != nil {
				break
			}
			if err = t.Execute(file, ctx); err != nil {
				break
			}
		} else {
			if _, err = file.Write(MustAsset(assetInfo.name)); err != nil {
				break
			}
		}

		if err = file.Close(); err != nil {
			break
		}
	}

	return err
}
