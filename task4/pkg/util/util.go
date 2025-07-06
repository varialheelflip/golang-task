package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetHeaderUserId(c *gin.Context) uint {
	userId, _ := c.Get("userID")
	parseUint, _ := strconv.ParseUint(fmt.Sprintf("%v", userId), 10, 64)
	return uint(parseUint)
}
