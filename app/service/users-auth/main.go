package main

import (
	"gitbus.com/exlab/zim-ms/app/service/users-auth/servants"
	"gitbus.com/exlab/zim-ms/app/service/users-auth/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zua "gitbus.com/exlab/zim-ms/app/service/users-auth/proto/gen/ZimUsersAuth"

	_ "gitbus.com/exlab/zim-ms/library/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	app := new(zua.UsersAuth)
	srv := servants.NewUsersAuth()
	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".AuthObj"
	app.AddServantWithContext(srv, obj)
	tars.Run()
}
