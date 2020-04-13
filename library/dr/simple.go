package dr

import (
	"net/http"

	"gitbus.com/exlab/zim-ms/library/errorx"
	"github.com/gin-gonic/gin"
)

type simpleServant struct {
	// just empty
}

type response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}

func (s *simpleServant) Success(c *gin.Context, data interface{}) {
	resp := &response{
		Code:   http.StatusOK,
		Msg:    "success",
		Result: data,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *simpleServant) Abort(c *gin.Context, err errorx.HttpError) {
	resp := &response{
		Code: err.Code(),
		Msg:  err.Error(),
	}
	c.JSON(err.Code(), resp)
}

func (s *simpleServant) Failure(c *gin.Context, err error) {
	resp := &response{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusInternalServerError, resp)
}

// NewSimpleServant return a simple BaseServant instance
func NewSimpleServant() BaseServant {
	return &simpleServant{}
}
