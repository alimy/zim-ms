package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/portal/sessions/internal/errorx"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/mirc/gen/api/sessions/manage"
	"github.com/gin-gonic/gin"

	zsm "gitbus.com/exlab/zim-ms/app/portal/sessions/proto/gen/ZimSessionsManage"
)

type manageSrv struct {
	baseServant
	manage *zsm.SessionsManage
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

// NewSessionsManage return a SessionsManage implement object
func NewSessionsManage() manage.SessionsManage {
	return &manageSrv{
		baseServant: newBaseServant(),
		manage:      newSessionsManageApp(),
	}
}
