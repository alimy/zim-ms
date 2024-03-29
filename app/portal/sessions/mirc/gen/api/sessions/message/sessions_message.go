// Code generated by go-mir. DO NOT EDIT.

package message

import (
	"github.com/gin-gonic/gin"
)

type SessionsMessage interface {
	Add(*gin.Context)
	Sub(*gin.Context)
}

// RegisterSessionsMessageServant register SessionsMessage servant to gin
func RegisterSessionsMessageServant(e *gin.Engine, s SessionsMessage) {
	router := e.Group("sessions/message")

	// register routes info to router
	router.Handle("GET", "/add/:lhs/:rhs", s.Add)
	router.Handle("GET", "/sub/:lhs/:rhs", s.Sub)
}
