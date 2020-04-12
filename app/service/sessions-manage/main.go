package main

import (
	"gitbus.com/exlab/zim-ms/app/service/sessions-manage/servants"
	"gitbus.com/exlab/zim-ms/app/service/sessions-manage/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zsm "gitbus.com/exlab/zim-ms/app/service/sessions-manage/proto/gen/ZimSessionsManage"

	_ "gitbus.com/exlab/zim-ms/app/service/sessions-manage/internal/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zsm.SessionsManage)
	srv := servants.NewSessionsManage()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".ManageObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
