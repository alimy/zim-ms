package main

import (
	"gitbus.com/exlab/zim-ms/app/admin/users/mirc/gen/api/admin/users/auth"
	"gitbus.com/exlab/zim-ms/app/admin/users/mirc/gen/api/admin/users/manage"
	"gitbus.com/exlab/zim-ms/app/admin/users/mirc/gen/api/admin/users/query"
	"gitbus.com/exlab/zim-ms/app/admin/users/servants"
	"gitbus.com/exlab/zim-ms/app/admin/users/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"

	_ "gitbus.com/exlab/zim-ms/library/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	e := gin.Default()
	registerServants(e)

	mux := &tars.TarsHttpMux{}
	mux.HandleFunc("/", e.ServeHTTP)

	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".AdminObj"
	tars.AddHttpServant(mux, obj)
	tars.Run()
}

func registerServants(e *gin.Engine) {
	auth.RegisterUsersAuthServant(e, servants.NewUsersAuth())
	query.RegisterUsersQueryServant(e, servants.NewUsersQuery())
	manage.RegisterUsersManageServant(e, servants.NewUsersManage())
}
