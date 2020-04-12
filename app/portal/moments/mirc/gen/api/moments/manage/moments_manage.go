// Code generated by go-mir. DO NOT EDIT.

package manage

import (
	"github.com/gin-gonic/gin"
)

type MomentsManage interface {
	Add(*gin.Context)
	Sub(*gin.Context)
}

// RegisterMomentsManageServant register MomentsManage servant to gin
func RegisterMomentsManageServant(e *gin.Engine, s MomentsManage) {
	router := e.Group("moments/manage")

	// register routes info to router
	router.Handle("GET", "/add/:lhs/:rhs", s.Add)
	router.Handle("GET", "/sub/:lhs/:rhs", s.Sub)
}