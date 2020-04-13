package main

import (
	"gitbus.com/exlab/zim-ms/app/service/sessions-query/servants"
	"gitbus.com/exlab/zim-ms/app/service/sessions-query/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zsq "gitbus.com/exlab/zim-ms/app/service/sessions-query/proto/gen/ZimSessionsQuery"

	_ "gitbus.com/exlab/zim-ms/library/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zsq.SessionsQuery)
	srv := servants.NewSessionsQuery()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".QueryObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
