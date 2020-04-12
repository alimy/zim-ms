package servants

import (
	"gitbus.com/exlab/zim-ms/app/admin/users/internal/locator"
	"gitbus.com/exlab/zim-ms/app/admin/users/internal/utils"

	zua "gitbus.com/exlab/zim-ms/app/admin/users/proto/gen/ZimUsersAuth"
	zum "gitbus.com/exlab/zim-ms/app/admin/users/proto/gen/ZimUsersManage"
	zuq "gitbus.com/exlab/zim-ms/app/admin/users/proto/gen/ZimUsersQuery"
)

func newUsersAuthApp() *zua.UsersAuth {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.UsersAuth.AuthObj")
	app := new(zua.UsersAuth)
	comm.StringToProxy(obj, app)
	return app
}

func newUsersQueryApp() *zuq.UsersQuery {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.UsersQuery.QueryObj")
	app := new(zuq.UsersQuery)
	comm.StringToProxy(obj, app)
	return app
}

func newUsersManageApp() *zum.UsersManage {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.UsersManage.ManageObj")
	app := new(zum.UsersManage)
	comm.StringToProxy(obj, app)
	return app
}
