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
