package routes

import (
	"squarecode/dynastymasra/rhino/api/handler"

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

func Index(router *gin.RouterGroup) {
	router.GET("/", handler.Index)
	router.GET("/ping", handler.Pong)
}
