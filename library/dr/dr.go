package dr

import (
	"gitbus.com/exlab/zim-ms/library/errorx"
	"github.com/gin-gonic/gin"
)

// BaseServant base servant interface
type BaseServant interface {
	Success(*gin.Context, interface{})
	Abort(*gin.Context, errorx.HttpError)
	Failure(*gin.Context, error)
}