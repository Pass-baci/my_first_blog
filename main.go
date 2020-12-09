package main

import (
	"ginblog/module"
	"ginblog/routes"
)

func main() {
	module.InitDataBase()
	routes.InitRouter()
}
