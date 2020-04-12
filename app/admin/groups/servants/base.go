package servants

import (
	"net/http"

	"gitbus.com/exlab/zim-ms/app/admin/groups/internal/errorx"
	"github.com/gin-gonic/gin"
)

type baseServant struct {
	// just empty
}

type response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}

func (s *baseServant) success(c *gin.Context, data interface{}) {
	resp := &response{
		Code:   http.StatusOK,
		Msg:    "success",
		Result: data,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *baseServant) abort(c *gin.Context, err errorx.HttpError) {
	resp := &response{
		Code: err.Code(),
		Msg:  err.Error(),
	}
	c.JSON(err.Code(), resp)
}

func (s *baseServant) failure(c *gin.Context, err error) {
	resp := &response{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func newBaseServant() baseServant {
	return baseServant{}
}
