package main

import (
	"log"
	"runtime"
	"squarecode/dynastymasra/rhino/api/routes"

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

func configuration() {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)
	log.Printf("Running with %d CPUs...\n", cpu)
}

func routers() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		routes.Person(v1)
		routes.Index(v1)
	}
	router.Run(":8000")
}

func main() {
	log.Println("Client is running...")
	configuration()
	routers()
}
