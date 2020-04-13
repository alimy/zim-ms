// Code generated by go-mir. DO NOT EDIT.

package query

import (
	"github.com/gin-gonic/gin"
)

type DnsQuery interface {
	Add(*gin.Context)
	Sub(*gin.Context)
}

// RegisterDnsQueryServant register DnsQuery servant to gin
func RegisterDnsQueryServant(e *gin.Engine, s DnsQuery) {
	router := e.Group("httpdns/query")

	// register routes info to router
	router.Handle("GET", "/add/:lhs/:rhs", s.Add)
	router.Handle("GET", "/sub/:lhs/:rhs", s.Sub)
}
