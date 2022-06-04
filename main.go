package main

import (
	"douyinProject/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repository.Init()
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
