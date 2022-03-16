package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    apiServer := gin.New()
    apiServer.LoadHTMLGlob("index.html")
    apiServer.Static("/css","css")
    apiServer.Static("/assets","assets")
    apiServer.Static("/js","js")
    apiServer.GET("/index", Index)

    apiServer.Run(":3388")
}

func Index(context *gin.Context) {
    context.HTML(http.StatusOK, "index.html", "")
}
