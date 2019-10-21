## 本项目分别演示
* go gin Web框架的使用
* go如何使用下面4种方式操作MySQL
    - database/sql操作MySQL
    - GORM:一款全功能的ORM go框架
    - go-xorm
    - GoMybatis
    
### 添加依赖
gin web框架: ```go get -u github.com/gin-gonic/gin```
MySQL连接库 database/sql: ```go get -u github.com/go-sql-driver/mysql```
GROM框架: ```go get -u github.com/jinzhu/gorm```

### 新建数据库
```sql
create schema go_study
```

### database/sql操作MySQL
使用原生database/sql操作MySQL

#### 如何做
1. 在上面数据库中新建数据表
```sql
CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(40) NOT NULL DEFAULT '',
  `last_name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

2. 修改main.go如下
```go
package main

import (
	db "webStudy/database"
)

func main() {
	defer db.SqlDB.Close()
	defer db.GromDb.Close()

	router := initRouter()
	router.Run(":8000")

	//router := gromInitRouter()
	//router.Run(":8000")
}
```

3. 执行命令
```shell script
go run *.go
```

4. 测试(GORM go-xorm GoMybatis的测试也类似)
首页
```shell script
curl http://127.0.0.1:8000
```
新增数据
```shell script
curl -X POST -F 'first_name=david' -F 'last_name=hong' http://127.0.0.1:8000/person
```
获取所有记录
```shell script
curl http://127.0.0.1:8000/persons
```
通过id获取一个记录
```shell script
# 获取id=1的记录
curl http://127.0.0.1:8000/person/1
```
通过id更新一个记录
```shell script
# 更新id=1的记录
curl -X PUT -H "Content-Type: application/json" -d '{"first_name":"123","last_name":"456"}' http://127.0.0.1:8000/person/1
# 或者
curl -X PUT -d 'first_name=123&last_name=456' http://127.0.0.1:8000/person/1
```
通过id删除一个记录
```shell script
# 删除id=1的记录
curl -X DELETE http://127.0.0.1:8000/person/1
```

### [GORM:一款全功能的ORM go框架](https://github.com/jinzhu/gorm) 
本项目中演示了最简单的CURD操作(apis/gromPerson.go)，更多操作内容可参考[文档](https://gorm.io/zh_CN/docs/index.html)
GROM有链式操作，使用了builder设计模式，和mybatis-plus一样方便易用

#### 如何演示GORM
修改main.go如下
```go
package main

import (
	db "webStudy/database"
)

func main() {
	defer db.SqlDB.Close()
	defer db.GromDb.Close()

	//router := initRouter()
	//router.Run(":8000")

	router := gromInitRouter()
	router.Run(":8000")
}
```
执行命令
```shell script
go run *.go
```

### 3. [go-xorm](https://github.com/go-xorm/xorm)
xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。

### 4. [GoMybatis](https://github.com/zhuxiujia/GoMybatis)
类似Mybatis的方式操作MySQL; 用go语言生成xml; 和Mybatis3.0无缝连接，方便Java项目转go项目