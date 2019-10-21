package main

import (
	"github.com/gin-gonic/gin"
	. "webStudy/apis"
)

/*
	路由: 演示database/sql操作MySQL
*/
func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/person", AddPersonApi)

	router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", UpdatePersonApi)

	router.DELETE("/person/:id", DeletePersonApi)

	return router
}
