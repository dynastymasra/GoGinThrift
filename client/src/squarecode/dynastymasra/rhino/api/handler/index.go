package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":   "Index",
		"status": "success",
	})
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   "PONG!",
	})
}
