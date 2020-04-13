package main

import (
	"context"
	"fmt"

	"gitbus.com/exlab/zim-ms/library/locator"
	"gitbus.com/exlab/zim-ms/library/utils"

	zdq "gitbus.com/exlab/zim-ms/app/service/httpdns/proto/gen/ZimDnsQuery"
	zdu "gitbus.com/exlab/zim-ms/app/service/httpdns/proto/gen/ZimDnsUpdate"
)

func main() {
	var res int32
	comm := locator.MyCommunicator()

	c := context.Background()
	rqx := map[string]string{
		"demo": "true",
	}
	queryObj := utils.ObjFrom("zim.HttpDns.QueryObj", "tcp -h 127.0.0.1 -p 10028 -t 60000")
	query := new(zdq.DnsQuery)
	comm.StringToProxy(queryObj, query)
	fmt.Printf("query reqest=> add 10 and 10  with reqCtx:%v\n", rqx)
	ret, err := query.AddWithContext(c, 10, 10, &res, rqx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("query response=> ret:%d result:%d reqCtx:%v\n", ret, res, rqx)

	c = context.Background()
	rqx = map[string]string{
		"demo": "true",
	}
	updateObj := utils.ObjFrom("zim.HttpDns.UpdateObj", "tcp -h 127.0.0.1 -p 10029 -t 60000")
	update := new(zdu.DnsUpdate)
	comm.StringToProxy(updateObj, update)
	fmt.Printf("update reqest=> add 10 and 10  with reqCtx:%v\n", rqx)
	ret, err = update.AddWithContext(c, 10, 10, &res, rqx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("update  response=> ret:%d result:%d reqCtx:%v\n", ret, res, rqx)
}
