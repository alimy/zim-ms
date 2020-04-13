package main

import (
	"gitbus.com/exlab/zim-ms/app/admin/httpdns/mirc/gen/api/admin/httpdns/query"
	"gitbus.com/exlab/zim-ms/app/admin/httpdns/mirc/gen/api/admin/httpdns/update"
	"gitbus.com/exlab/zim-ms/app/admin/httpdns/servants"
	"gitbus.com/exlab/zim-ms/app/admin/httpdns/version"
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
	query.RegisterDnsQueryServant(e, servants.NewDnsQuery())
	update.RegisterDnsUpdateServant(e, servants.NewDnsUpdate())
}
