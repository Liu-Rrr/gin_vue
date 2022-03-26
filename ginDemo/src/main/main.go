package main

import (
	"ginDemo/src/main/dbUtils"
	"ginDemo/src/main/router"
)

func main() {

	dbUtils.Init()
	r := router.InitRouter()

	r.Run(":9292")
}
