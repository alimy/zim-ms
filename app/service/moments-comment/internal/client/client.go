package main

import (
	"context"
	"fmt"

	"gitbus.com/exlab/zim-ms/library/locator"
	"gitbus.com/exlab/zim-ms/library/utils"

	zmc "gitbus.com/exlab/zim-ms/app/service/moments-manage/proto/gen/ZimMomentsComment"
)

func main() {
	comm := locator.MyCommunicator()
	obj := utils.ObjFrom("zim.MomentsComment.CommentObj", "tcp -h 127.0.0.1 -p 10028 -t 60000")
	app := new(zmc.MomentsComment)
	comm.StringToProxy(obj, app)

	var res int32
	c := context.Background()
	rqx := map[string]string{
		"demo": "true",
	}

	fmt.Printf("reqest=> add 10 and 10  with reqCtx:%v\n", rqx)
	ret, err := app.AddWithContext(c, 10, 10, &res, rqx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response=> ret:%d result:%d reqCtx:%v\n", ret, res, rqx)
}
