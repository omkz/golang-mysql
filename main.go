package main

import (
	"github.com/omkz/golang-mysql/config"
	"github.com/omkz/golang-mysql/controllers"
)


func init(){
	config.ConnectDB("root:root@tcp(127.0.0.1:3306)/golang_blog")
}

func main() {
	controllers.GetPosts()
}