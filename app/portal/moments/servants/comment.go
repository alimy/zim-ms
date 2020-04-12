package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/portal/moments/internal/errorx"
	"gitbus.com/exlab/zim-ms/app/portal/moments/mirc/gen/api/moments/comment"
	"github.com/gin-gonic/gin"

	zmc "gitbus.com/exlab/zim-ms/app/portal/moments/proto/gen/ZimMomentsComment"
)

type commentSrv struct {
	baseServant
	comment *zmc.MomentsComment
}

func (s *commentSrv) Add(c *gin.Context) {
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
	if _, err := s.comment.Add(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

func (s *commentSrv) Sub(c *gin.Context) {
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
	if _, err := s.comment.Sub(lh, rh, &res); err == nil {
		s.success(c, res)
	} else {
		s.failure(c, err)
	}
}

// NewMomentsComment return a MomentsComment implement object
func NewMomentsComment() comment.MomentsComment {
	return &commentSrv{
		baseServant: newBaseServant(),
		comment:     newMomentsCommentApp(),
	}
}
