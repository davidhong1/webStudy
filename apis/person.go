package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	. "webStudy/models"
)

// 首页，单独测试web是否可用 url=http://127.0.0.1/
func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "演示 MySQL首页")
}

// 新建person记录，post请求，提交form表单
func AddPersonApi(c *gin.Context) {
	//获取参数
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}
	//添加记录
	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// 获取所有person记录
func GetPersonsApi(c *gin.Context) {
	p := Person{}
	//获取所有记录
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, persons)
}

// 通过restful方式提交id，获取person记录 get请求
func GetPersonApi(c *gin.Context) {
	//获取参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln()
	}

	p := Person{Id: id}
	person, err := p.GetPerson()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "找不到该记录",
		})
		return
	}
	c.JSON(http.StatusOK, person)
}

// 通过restful方式使用put请求更新记录，同时允许form表单和json格式body
func UpdatePersonApi(c *gin.Context) {
	//获取参数
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	person := Person{Id: id}
	//绑定form表单或json到person对象
	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := person.UpdatePerson()

	msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// 通过id删除记录
func DeletePersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}
	ra, err := p.DeletePerson()

	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
