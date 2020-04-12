// Code generated by go-mir. DO NOT EDIT.

package query

import (
	"github.com/gin-gonic/gin"
)

type SessionsQuery interface {
	Add(*gin.Context)
	Sub(*gin.Context)
}

// RegisterSessionsQueryServant register SessionsQuery servant to gin
func RegisterSessionsQueryServant(e *gin.Engine, s SessionsQuery) {
	router := e.Group("sessions/query")

	// register routes info to router
	router.Handle("GET", "/add/:lhs/:rhs", s.Add)
	router.Handle("GET", "/sub/:lhs/:rhs", s.Sub)
}