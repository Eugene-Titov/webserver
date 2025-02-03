package main

import (
	"fmt"
	"ip/ip"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	fmt.Println("index--------------------")
	fmt.Println("index--------------------end")
}

var ipServer = ip.CreatePortServer(8080)

func main() {
	router := gin.Default()
	router.GET("/", index)
	// router.GET("/registration", registration)

	fmt.Printf("start web-server on port %s\n", ipServer.GetPort())
	router.Run()
	fmt.Println("exit")
}
