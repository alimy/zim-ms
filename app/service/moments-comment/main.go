package main

import (
	"gitbus.com/exlab/zim-ms/library/logus"
	"gitbus.com/exlab/zim-ss/app/service/moments-manage/servants"
	"gitbus.com/exlab/zim-ss/app/service/moments-manage/version"
	"github.com/TarsCloud/TarsGo/tars"

	zmc "gitbus.com/exlab/zim-ss/app/service/moments-manage/proto/gen/ZimMomentsComment"

	_ "gitbus.com/exlab/zim-ss/app/service/moments-manage/internal/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zmc.MomentsComment)
	srv := servants.NewMomentsComment()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".CommentObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
