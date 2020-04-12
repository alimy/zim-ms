package servants

import (
	"gitbus.com/exlab/zim-ms/app/portal/sessions/internal/locator"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/internal/utils"

	zsm "gitbus.com/exlab/zim-ms/app/portal/sessions/proto/gen/ZimSessionsManage"
	zse "gitbus.com/exlab/zim-ms/app/portal/sessions/proto/gen/ZimSessionsMessage"
	zsq "gitbus.com/exlab/zim-ms/app/portal/sessions/proto/gen/ZimSessionsQuery"
)

func newSessionsQueryApp() *zsq.SessionsQuery {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.SessionsQuery.QueryObj")
	app := new(zsq.SessionsQuery)
	comm.StringToProxy(obj, app)
	return app
}

func newSessionsManageApp() *zsm.SessionsManage {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.SessionsManage.ManageObj")
	app := new(zsm.SessionsManage)
	comm.StringToProxy(obj, app)
	return app
}

func newSessionsMessageApp() *zse.SessionsMessage {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.SessionsMessage.MessageObj")
	app := new(zse.SessionsMessage)
	comm.StringToProxy(obj, app)
	return app
}
