package servants

import (
	"gitbus.com/exlab/zim-ms/library/locator"
	"gitbus.com/exlab/zim-ms/library/utils"

	zdq "gitbus.com/exlab/zim-ms/app/admin/httpdns/proto/gen/ZimDnsQuery"
	zdu "gitbus.com/exlab/zim-ms/app/admin/httpdns/proto/gen/ZimDnsUpdate"
)

func newDnsQueryApp() *zdq.DnsQuery {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.HttpDns.QueryObj")
	app := new(zdq.DnsQuery)
	comm.StringToProxy(obj, app)
	return app
}

func newDnsUpdateApp() *zdu.DnsUpdate {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.HttpDns.UpdateObj")
	app := new(zdu.DnsUpdate)
	comm.StringToProxy(obj, app)
	return app
}
