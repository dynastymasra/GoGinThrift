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

func Person(router *gin.RouterGroup) {
	router.GET("/persons", handler.GetAll)
	router.GET("/person/:id", handler.GetById)
	router.POST("/person", handler.Create)
	router.PUT("/person", handler.Update)
	router.DELETE("/person", handler.Delete)
}
