package servants

import (
	"gitbus.com/exlab/zim-ms/library/locator"
	"gitbus.com/exlab/zim-ms/library/utils"

	zdq "gitbus.com/exlab/zim-ms/app/portal/httpdns/proto/gen/ZimDnsQuery"
)

func newDnsQueryApp() *zdq.DnsQuery {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.HttpDns.QueryObj")
	app := new(zdq.DnsQuery)
	comm.StringToProxy(obj, app)
	return app
}
