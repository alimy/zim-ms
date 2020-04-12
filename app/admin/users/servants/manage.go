package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/admin/users/internal/errorx"
	"gitbus.com/exlab/zim-ms/app/admin/users/mirc/gen/api/admin/users/manage"
	"github.com/gin-gonic/gin"

	zum "gitbus.com/exlab/zim-ms/app/admin/users/proto/gen/ZimUsersManage"
)

type manageSrv struct {
	baseServant

	manage *zum.UsersManage
}

func (s *manageSrv) Add(c *gin.Context) {
	lhs, rhs := c.Param("lhs"), c.Param("rhs")
	var (
		lh, rh, res int32
		v           int
		err1, err2  error
	)

	if v, err1 = strconv.Atoi(lhs); err1 == nil {
		lh = int32(v)
	}
	if v, err2 = strconv.Atoi(rhs); err2 == nil {
		rh = int32(v)
	}
	if err1 != nil || err2 != nil {
		s.abort(c, errorx.ErrParamNotValide)
		return
	}
	if _, err := s.manage.Add(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

func (s *manageSrv) Sub(c *gin.Context) {
	lhs, rhs := c.Param("lhs"), c.Param("rhs")
	var (
		lh, rh, res int32
		v           int
		err1, err2  error
	)

	if v, err1 = strconv.Atoi(lhs); err1 == nil {
		lh = int32(v)
	}
	if v, err2 = strconv.Atoi(rhs); err2 == nil {
		rh = int32(v)
	}
	if err1 != nil || err2 != nil {
		s.abort(c, errorx.ErrParamNotValide)
		return
	}
	if _, err := s.manage.Sub(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

// NewUsersManage return a UsersManage implement object
func NewUsersManage() manage.UsersManage {
	return &manageSrv{
		baseServant: newBaseServant(),
		manage:      newUsersManageApp(),
	}
}
