package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(context *gin.Context, data interface{}, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  msg,
	})
}
func Success(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
}
func Err(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}
