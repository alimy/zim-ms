// Code generated by go-mir. DO NOT EDIT.

package manage

import (
	"github.com/gin-gonic/gin"
)

type UsersManage interface {
	Add(*gin.Context)
	Sub(*gin.Context)
}

// RegisterUsersManageServant register UsersManage servant to gin
func RegisterUsersManageServant(e *gin.Engine, s UsersManage) {
	router := e.Group("admin/users/manage")

	// register routes info to router
	router.Handle("GET", "/add/:lhs/:rhs", s.Add)
	router.Handle("GET", "/sub/:lhs/:rhs", s.Sub)
}