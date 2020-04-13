package servants

import (
	"strconv"

	"gitbus.com/exlab/zim-ms/app/portal/moments/mirc/gen/api/moments/comment"
	"gitbus.com/exlab/zim-ms/library/dr"
	"gitbus.com/exlab/zim-ms/library/errorx"
	"github.com/gin-gonic/gin"

	zmc "gitbus.com/exlab/zim-ms/app/portal/moments/proto/gen/ZimMomentsComment"
)

type commentSrv struct {
	dr.BaseServant
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
		s.Abort(c, errorx.ErrParamNotValide)
		return
	}
	if _, err := s.comment.Add(lh, rh, &res); err == nil {
		s.Success(c, res)
	} else {
		s.Failure(c, err)
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
		s.Abort(c, errorx.ErrParamNotValide)
		return
	}
	if _, err := s.comment.Sub(lh, rh, &res); err == nil {
		s.Success(c, res)
	} else {
		s.Failure(c, err)
	}
}

// NewMomentsComment return a MomentsComment implement object
func NewMomentsComment() comment.MomentsComment {
	return &commentSrv{
		BaseServant: dr.NewSimpleServant(),
		comment:     newMomentsCommentApp(),
	}
}
