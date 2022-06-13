package middleware

import (
	JwtLib "douyin/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		fmt.Printf("PATH: %v | USE TIME: %v | RESPONSE STATUS: %d\n", c.Request.Method+c.FullPath(), latency, c.Writer.Status())
	}
}
func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")

		_, err := JwtLib.JwtParse(token)
		if err != nil {
			log.Println("中间件报错:", err)
			c.JSON(http.StatusBadGateway, gin.H{})
			return
		} else {
			c.Next()
			log.Println("token 通过")

		}
	}
}
