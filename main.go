package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snake/middlewares"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.Cors())

	r.StaticFS("/web", http.Dir("./web"))

	// 监听80端口
	r.Run("0.0.0.0:80")

}
