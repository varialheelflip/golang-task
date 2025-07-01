package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一API响应格式
func Response(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
		"data":    data,
	})
}

func Success(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, "success", data)
}

func BadRequest(c *gin.Context, message string) {
	Response(c, http.StatusBadRequest, message, nil)
}

func ServerError(c *gin.Context, message string) {
	Response(c, http.StatusInternalServerError, message, nil)
}
