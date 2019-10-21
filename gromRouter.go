package main

import (
	"github.com/gin-gonic/gin"
	. "webStudy/apis"
)

/*
	演示GROM操作数据库
*/
func gromInitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", GromIndexApi)

	router.POST("/person", GromAddPersonApi)

	router.GET("/persons", GromGetPersonsApi)

	router.GET("/person/:id", GromGetPersonApi)

	router.PUT("/person/:id", GromUpdatePersonApi)

	router.DELETE("/person/:id", GromDeletePersonApi)

	return router
}
