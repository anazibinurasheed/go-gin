package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ok := false
	i := 0
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		if !ok {
			ok = true
			time.Sleep(5 * time.Second) // 5 Seconds
		}
		fmt.Println(ok)
		i++
		c.JSON(200, gin.H{
			"message": "pong" + fmt.Sprint(i),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
