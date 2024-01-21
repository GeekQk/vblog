package response

import (
	"net/http"

	"github.com/GeekQk/vblog/exception"
	"github.com/gin-gonic/gin"
)

// 响应成功 返回数据
func Success(c *gin.Context, data any) {

}

// 响应失败 返回Api Exception
func Failed(c *gin.Context, err error) {
	//先返回异常
	var resp *exception.APIExeption
	if e, ok := err.(*exception.APIExeption); ok {
		resp = e
	} else {
		resp = exception.NewApiException(500, http.StatusText(http.StatusInternalServerError)).
			WithMessage(err.Error()).WithHttpCode(500)
	}
	// 返回响应
	c.JSON(resp.HttpCode, resp)
	//中断
	c.Abort()

}
