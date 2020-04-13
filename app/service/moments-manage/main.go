package main

import (
	"gitbus.com/exlab/zim-ms/app/service/moments-manage/servants"
	"gitbus.com/exlab/zim-ms/app/service/moments-manage/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zmm "gitbus.com/exlab/zim-ms/app/service/moments-manage/proto/gen/ZimMomentsManage"

	_ "gitbus.com/exlab/zim-ms/library/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zmm.MomentsManage)
	srv := servants.NewMomentsManage()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".ManageObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
