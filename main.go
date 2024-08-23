package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    // Загрузка HTML-шаблонов из директории templates
    router.LoadHTMLGlob("templates/*")

    // Обслуживание статических файлов из директории static
    router.Static("/static", "./static")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title":   "Welcome to Gin",
        })
    })

    // Обслуживание видео
    router.GET("/video", func(c *gin.Context) {
        c.File("./static/video.mp4") // Путь к вашему видеофайлу
    })

    router.Run(":9200")
}
