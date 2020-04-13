package main

import (
	"gitbus.com/exlab/zim-ms/app/service/users-manage/servants"
	"gitbus.com/exlab/zim-ms/app/service/users-manage/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zum "gitbus.com/exlab/zim-ms/app/service/users-manage/proto/gen/ZimUsersManage"

	_ "gitbus.com/exlab/zim-ms/library/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zum.UsersManage)
	srv := servants.NewUsersManage()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".ManageObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
