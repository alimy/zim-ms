package main

import (
	"gitbus.com/exlab/zim-ms/app/service/sessions-message/servants"
	"gitbus.com/exlab/zim-ms/app/service/sessions-message/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zsm "gitbus.com/exlab/zim-ms/app/service/sessions-message/proto/gen/ZimSessionsMessage"

	_ "gitbus.com/exlab/zim-ms/app/service/sessions-message/internal/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zsm.SessionsMessage)
	srv := servants.NewSessionsMessage()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".MessageObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
