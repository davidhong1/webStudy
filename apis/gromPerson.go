package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	db "webStudy/database"
	. "webStudy/models"
)

// 首页，单独测试web是否可用 url=http://127.0.0.1/
func GromIndexApi(c *gin.Context) {
	c.String(http.StatusOK, "演示grom index")
}

// 新建person记录，post请求，提交form表单
func GromAddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := GromPerson{FirstName: firstName, LastName: lastName}

	//没有该表时，新建表
	db.GromDb.AutoMigrate(&GromPerson{})
	//插入p对象
	db.GromDb.Create(&p)

	c.JSON(http.StatusOK, gin.H{
		"msg": "插入一个GromPerson",
	})
}

// 获取所有person记录
func GromGetPersonsApi(c *gin.Context) {
	var persons []GromPerson
	db.GromDb.Find(&persons)
	c.JSON(http.StatusOK, persons)
}

// 通过restful方式提交id，获取person记录 get请求
func GromGetPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln()
	}
	var person GromPerson
	//通过主键获取记录
	db.GromDb.First(&person, id)
	c.JSON(http.StatusOK, person)
}

// 通过restful方式使用put请求更新记录，同时允许form表单和json格式body
func GromUpdatePersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	person := GromPerson{Id: id}
	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}
	//若没有该ID的记录，则新建该ID的记录
	db.GromDb.Save(&person)

	msg := fmt.Sprintf("Update person %d successful", person.Id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// 通过id删除记录
func GromDeletePersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := GromPerson{Id: id}
	db.GromDb.Delete(&p)

	msg := fmt.Sprintf("Delete person %d successful", p.Id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
