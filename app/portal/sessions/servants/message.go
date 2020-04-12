package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/portal/sessions/internal/errorx"
	"gitbus.com/exlab/zim-ms/app/portal/sessions/mirc/gen/api/sessions/message"
	"github.com/gin-gonic/gin"

	zsm "gitbus.com/exlab/zim-ms/app/portal/sessions/proto/gen/ZimSessionsMessage"
)

type messageSrv struct {
	baseServant
	message *zsm.SessionsMessage
}

func (s *messageSrv) Add(c *gin.Context) {
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
	if _, err := s.message.Add(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

func (s *messageSrv) Sub(c *gin.Context) {
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
	if _, err := s.message.Sub(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

// NewSessionsMessage return a SessionsMessage implement object
func NewSessionsMessage() message.SessionsMessage {
	return &messageSrv{
		baseServant: newBaseServant(),
		message:     newSessionsMessageApp(),
	}
}
