package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/admin/groups/internal/errorx"
	"gitbus.com/exlab/zim-ms/app/admin/groups/mirc/gen/api/admin/groups/auth"
	"github.com/gin-gonic/gin"

	zua "gitbus.com/exlab/zim-ms/app/admin/groups/proto/gen/ZimUsersAuth"
)

type authSrv struct {
	baseServant
	auth *zua.UsersAuth
}

func (s *authSrv) Add(c *gin.Context) {
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
	if _, err := s.auth.Add(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

func (s *authSrv) Sub(c *gin.Context) {
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
	if _, err := s.auth.Sub(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

// NewUsersAuth return a UsersAuth implement object
func NewUsersAuth() auth.UsersAuth {
	return &authSrv{
		baseServant: newBaseServant(),
		auth:        newUsersAuthApp(),
	}
}
