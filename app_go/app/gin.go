package main

import "github.com/gin-gonic/gin"

func main() {
        r := gin.Default()
        r.GET("/", func(c *gin.Context) {
                c.String(200, "Hello,World!")
        })
        r.GET("/hoge", func(c *gin.Context){
                c.String(200,"Hoge!!!!!")
        })
        r.Run(":8001")
    }
