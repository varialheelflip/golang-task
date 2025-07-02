package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"unicode"
)

func StrNotBlank(s string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return !unicode.IsSpace(r)
	}) >= 0
}

func GetHeaderUserId(c *gin.Context) uint {
	userId, _ := c.Get("userID")
	parseUint, _ := strconv.ParseUint(fmt.Sprintf("%v", userId), 10, 64)
	return uint(parseUint)
}
