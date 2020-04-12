package main

import (
	"gitbus.com/exlab/zim-ms/app/service/users-query/servants"
	"gitbus.com/exlab/zim-ms/app/service/users-query/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zuq "gitbus.com/exlab/zim-ms/app/service/users-query/proto/gen/ZimUsersQuery"

	_ "gitbus.com/exlab/zim-ms/app/service/users-query/internal/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zuq.UsersQuery)
	srv := servants.NewUsersQuery()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".QueryObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
