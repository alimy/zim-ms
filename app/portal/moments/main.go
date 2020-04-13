package main

import (
	"gitbus.com/exlab/zim-ms/app/portal/moments/mirc/gen/api/moments/comment"
	"gitbus.com/exlab/zim-ms/app/portal/moments/mirc/gen/api/moments/manage"
	"gitbus.com/exlab/zim-ms/app/portal/moments/mirc/gen/api/moments/query"
	"gitbus.com/exlab/zim-ms/app/portal/moments/servants"
	"gitbus.com/exlab/zim-ms/app/portal/moments/version"
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
	query.RegisterMomentsQueryServant(e, servants.NewMomentsQuery())
	manage.RegisterMomentsManageServant(e, servants.NewMomentsManage())
	comment.RegisterMomentsCommentServant(e, servants.NewMomentsComment())
}
