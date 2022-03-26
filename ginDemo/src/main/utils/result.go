package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    "0",
	})
}
func Failed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "failed",
		"code":    "-1",
	})
}
func FailedErr(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": "failed",
		"code":    "-1",
		"err":     v,
	})
}
