package main

import (
	"todo-lists/external"
	"todo-lists/external/mysql"
)

func main() {
	defer mysql.CloseConn()

	external.Router.Run()
}