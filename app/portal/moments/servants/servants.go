package servants

import (
	"gitbus.com/exlab/zim-ms/library/locator"
	"gitbus.com/exlab/zim-ms/library/utils"

	zmc "gitbus.com/exlab/zim-ms/app/portal/moments/proto/gen/ZimMomentsComment"
	zmm "gitbus.com/exlab/zim-ms/app/portal/moments/proto/gen/ZimMomentsManage"
	zmq "gitbus.com/exlab/zim-ms/app/portal/moments/proto/gen/ZimMomentsQuery"
)

func newMomentsCommentApp() *zmc.MomentsComment {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.MomentsComment.CommentObj")
	app := new(zmc.MomentsComment)
	comm.StringToProxy(obj, app)
	return app
}

func newMomentsQueryApp() *zmq.MomentsQuery {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.MomentsQuery.QueryObj")
	app := new(zmq.MomentsQuery)
	comm.StringToProxy(obj, app)
	return app
}

func newMomentsManageApp() *zmm.MomentsManage {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.MomentsManage.ManageObj")
	app := new(zmm.MomentsManage)
	comm.StringToProxy(obj, app)
	return app
}
