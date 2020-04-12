// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"sort"
	"strings"
)

//go:generate go-bindata -nomemcopy -pkg=${GOPACKAGE} -ignore=README.md -prefix=templates -debug=false -o=templates_gen.go templates/...

// tmplCtx template context for generate project
type tmplCtx struct {
	AppName    string
	Category   string
	Target     string
	PkgName    string
	PacketName string
}

// tmplInfo template data info
type tmplInfo struct {
	name   string
	isTmpl bool
	isExec bool
}

type tmplInfos = map[string]tmplInfo

var (
	// tmplFiles map of generated file's assets info
	tmplFiles = make(map[string]tmplInfos, 2)

	// styles map of style slice to style
	styles = make(map[string]string)
)

// ts style slice alice type
type ts []string

func init() {
	for _, s := range []struct {
		styles ts
		target string
	}{
		{ts{"tars"}, "tars"},
		{ts{"mir"}, "tars/mir"},
		{ts{"tars", "mir"}, "tars/mir"},
	} {
		styles[s.styles.String()] = s.target
	}
}

func (s ts) String() string {
	sort.Strings(s)
	return strings.Join(s, ":")
}

func tmplInfosFrom(s []string) (tmplInfos, bool) {
	sort.Strings(s)
	style := styles[strings.Join(s, ":")]
	res, exist := tmplFiles[style]
	return res, exist
}
