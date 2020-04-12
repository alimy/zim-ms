package main

import (
	"gitbus.com/exlab/zim-ms/library/logus"
	"gitbus.com/exlab/zim-ss/app/service/moments-query/servants"
	"gitbus.com/exlab/zim-ss/app/service/moments-query/version"
	"github.com/TarsCloud/TarsGo/tars"

	zmq "gitbus.com/exlab/zim-ss/app/service/moments-query/proto/gen/ZimMomentsQuery"

	_ "gitbus.com/exlab/zim-ss/app/service/moments-query/internal/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zmq.MomentsQuery)
	srv := servants.NewMomentsQuery()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".QueryObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
