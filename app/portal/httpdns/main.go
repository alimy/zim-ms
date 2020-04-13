package main

import (
	"gitbus.com/exlab/zim-ms/app/portal/httpdns/mirc/gen/api/httpdns/query"
	"gitbus.com/exlab/zim-ms/app/portal/httpdns/servants"
	"gitbus.com/exlab/zim-ms/app/portal/httpdns/version"
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
	query.RegisterDnsQueryServant(e, servants.NewDnsQuery())

	mux := &tars.TarsHttpMux{}
	mux.HandleFunc("/", e.ServeHTTP)

	cfg := tars.GetServerConfig()
	obj := cfg.App + "." + cfg.Server + ".PortalObj"
	tars.AddHttpServant(mux, obj)
	tars.Run()
}
