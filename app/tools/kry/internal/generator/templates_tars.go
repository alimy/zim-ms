// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

func init() {
	tmplFiles["tars"] = tmplInfos{
		"version/version.go":                    {"tars/simple/version/version.tmpl", false, false},
		"servants/servants.go":                  {"tars/simple/servants/servants.tmpl", false, false},
		"go.mod":                                {"tars/simple/go_mod.tmpl", true, false},
		"proto/calculator.tars":                 {"tars/simple/proto/calculator_tars.tmpl", false, false},
		"proto/gen/calculator/Calculator_IF.go": {"tars/simple/proto/gen/calculator/Calculator_IF.tmpl", false, false},
		"service.conf":                          {"tars/simple/service_conf.tmpl", true, false},
		"makefile":                              {"tars/simple/makefile.tmpl", true, false},
		"internal/client/client.go":             {"tars/simple/internal/client/client.tmpl", true, false},
		"internal/internal.go":                  {"tars/simple/internal/internal.tmpl", false, false},
		"README.md":                             {"tars/simple/README.tmpl", false, false},
		".gitignore":                            {"tars/simple/gitignore.tmpl", true, false},
		"main.go":                               {"tars/simple/main.tmpl", true, false},
	}
	tmplFiles["tars/mir"] = tmplInfos{
		"version/version.go":                    {"tars/mir/version/version.tmpl", false, false},
		"servants/servants.go":                  {"tars/mir/servants/servants.tmpl", true, false},
		"servants/calculator.go":                {"tars/mir/servants/calculator.tmpl", true, false},
		"go.mod":                                {"tars/mir/go_mod.tmpl", true, false},
		"proto/calculator.tars":                 {"tars/mir/proto/calculator_tars.tmpl", false, false},
		"proto/gen/calculator/Calculator_IF.go": {"tars/mir/proto/gen/calculator/Calculator_IF.tmpl", false, false},
		"service.conf":                          {"tars/mir/service_conf.tmpl", true, false},
		"makefile":                              {"tars/mir/makefile.tmpl", true, false},
		"README.md":                             {"tars/mir/README.tmpl", false, false},
		".gitignore":                            {"tars/mir/gitignore.tmpl", true, false},
		"main.go":                               {"tars/mir/main.tmpl", true, false},
		"mirc/gen/api/calculator.go":            {"tars/mir/mirc/gen/api/calculator.tmpl", false, false},
		"mirc/routes/calculator.go":             {"tars/mir/mirc/routes/calculator.tmpl", false, false},
		"mirc/main.go":                          {"tars/mir/mirc/main.tmpl", true, false},
	}
}
