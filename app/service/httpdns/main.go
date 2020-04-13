package main

import (
	"gitbus.com/exlab/zim-ms/app/service/httpdns/servants"
	"gitbus.com/exlab/zim-ms/app/service/httpdns/version"
	"gitbus.com/exlab/zim-ms/library/logus"
	"github.com/TarsCloud/TarsGo/tars"

	zdq "gitbus.com/exlab/zim-ms/app/service/httpdns/proto/gen/ZimDnsQuery"
	zdu "gitbus.com/exlab/zim-ms/app/service/httpdns/proto/gen/ZimDnsUpdate"

	_ "gitbus.com/exlab/zim-ms/library/debug"
)

func init() {
	logus.SetLevel(logus.LevelDebug)
}

func main() {
	version.ShowInfo()

	cfg := tars.GetServerConfig()

	// add  DnsQuery servant
	queryApp := new(zdq.DnsQuery)
	querySrv := servants.NewDnsQuery()
	queryObj := cfg.App + "." + cfg.Server + ".QueryObj"
	queryApp.AddServantWithContext(querySrv, queryObj)

	// add  DnsUpdate servant
	updateApp := new(zdu.DnsUpdate)
	updateSrv := servants.NewDnsQuery()
	updateObj := cfg.App + "." + cfg.Server + ".UpdateObj"
	updateApp.AddServantWithContext(updateSrv, updateObj)

	tars.Run()
}
