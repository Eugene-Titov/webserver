package main

import (
	"fmt"
	"ip/ip"
	"pages/pages"

	"github.com/gin-gonic/gin"
)

var ipServer = ip.CreatePortServer(8080)
var pagesHTML = pages.CreatePager()

func index(c *gin.Context) {
	fmt.Println("index--------------------")
	html := pagesHTML.GetPage("index")
	fmt.Println(html)
	fmt.Println("index--------------------end")
}

func main() {
	router := gin.Default()
	router.GET("/", index)

	fmt.Printf("start web-server on port %s\n", ipServer.GetPort())
	router.Run()
	fmt.Println("exit")
}
