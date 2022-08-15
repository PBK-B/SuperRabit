package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(CorsMiddleware())

	r.Static("/", "/volume/common/view/index")
	port := os.Getenv("INDEX_WEB_PORT")
	if len(port) == 0 {
		port = "8601" //默认8601端口
	}
	log.Printf("🚀 App is running at port: " + port)
	http.ListenAndServe(":"+port, r)
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*") //Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*") //
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
