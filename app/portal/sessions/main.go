package main

import (
	"gitbus.com/exlab/zim-ms/app/portal/sessions/mirc/gen/api/sessions/manage"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/mirc/gen/api/sessions/message"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/mirc/gen/api/sessions/query"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/servants"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"

	_ "gitbus.com/exlab/zim-ms/app/portal/sessions/internal/debug"
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
	query.RegisterSessionsQueryServant(e, servants.NewSessionsQuery())
	manage.RegisterSessionsManageServant(e, servants.NewSessionsManage())
	message.RegisterSessionsMessageServant(e, servants.NewSessionsMessage())
}
