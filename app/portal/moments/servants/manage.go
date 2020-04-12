package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/portal/moments/internal/errorx"
	"gitbus.com/exlab/zim-ms/app/portal/moments/mirc/gen/api/moments/manage"
	"github.com/gin-gonic/gin"

	zmm "gitbus.com/exlab/zim-ms/app/portal/moments/proto/gen/ZimMomentsManage"
)

type manageSrv struct {
	baseServant
	manage *zmm.MomentsManage
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

// NewMomentsManage return a MomentsManage implement object
func NewMomentsManage() manage.MomentsManage {
	return &manageSrv{
		baseServant: newBaseServant(),
		manage:      newMomentsManageApp(),
	}
}
