package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/portal/sessions/internal/errorx"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/mirc/gen/api/sessions/query"
	"github.com/gin-gonic/gin"

	zsq "gitbus.com/exlab/zim-ms/app/portal/sessions/proto/gen/ZimSessionsQuery"
)

type querySrv struct {
	baseServant
	query *zsq.SessionsQuery
}

func (s *querySrv) Add(c *gin.Context) {
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
	if _, err := s.query.Add(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

func (s *querySrv) Sub(c *gin.Context) {
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
	if _, err := s.query.Sub(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

// NewSessionsQuery return a SessionsQuery implement object
func NewSessionsQuery() query.SessionsQuery {
	return &querySrv{
		baseServant: newBaseServant(),
		query:       newSessionsQueryApp(),
	}
}
