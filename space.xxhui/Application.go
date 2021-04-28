package main

import "github.com/gin-gonic/gin"

func main() {
	//router := gin.Default()//这种是声明变量并赋值
	var router *gin.Engine //声明变量
	router = gin.Default() //赋值

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
