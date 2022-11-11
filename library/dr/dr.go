package dr

import (
	"gitbus.com/exlab/zim-ms/library/errorx"
	"github.com/gin-gonic/gin"
)

// BaseServant base servant interface
type BaseServant interface {
	Success(*gin.Context, any)
	Abort(*gin.Context, errorx.HttpError)
	Failure(*gin.Context, error)
}
