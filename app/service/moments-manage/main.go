package main

import (
	"gitbus.com/exlab/zim-ms/library/logus"
	"gitbus.com/exlab/zim-ss/app/service/moments-manage/servants"
	"gitbus.com/exlab/zim-ss/app/service/moments-manage/version"
	"github.com/TarsCloud/TarsGo/tars"

	zmm "gitbus.com/exlab/zim-ss/app/service/moments-manage/proto/gen/ZimMomentsManage"

	_ "gitbus.com/exlab/zim-ss/app/service/moments-manage/internal/debug"
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
