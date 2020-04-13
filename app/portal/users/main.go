package main

import (
	"gitbus.com/exlab/zim-ms/app/portal/users/mirc/gen/api/users/auth"
	"gitbus.com/exlab/zim-ms/app/portal/users/mirc/gen/api/users/manage"
	"gitbus.com/exlab/zim-ms/app/portal/users/mirc/gen/api/users/query"
	"gitbus.com/exlab/zim-ms/app/portal/users/servants"
	"gitbus.com/exlab/zim-ms/app/portal/users/version"
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
	obj := cfg.App + "." + cfg.Server + ".PortalObj"
	tars.AddHttpServant(mux, obj)
	tars.Run()
}

func registerServants(e *gin.Engine) {
	auth.RegisterUsersAuthServant(e, servants.NewUsersAuth())
	query.RegisterUsersQueryServant(e, servants.NewUsersQuery())
	manage.RegisterUsersManageServant(e, servants.NewUsersManage())
}
