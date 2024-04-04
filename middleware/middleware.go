package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	fmt.Println("Inside middleware")
	return func(c *gin.Context) {
		start := time.Now()

		// Read request body
		reqBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
		}

		// Restore request body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		// Process request
		c.Next()

		// Log request and response details
		fmt.Println("---------------------------------------------------------------")
		fmt.Printf("[%s] %s %s | Request: %s | Response: %d %s | Duration: %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			string(reqBody),
			c.Writer.Status(),
			http.StatusText(c.Writer.Status()),
			time.Since(start),
		)
		fmt.Println("---------------------------------------------------------------")
	}
}

// func main() {
// 	router := gin.Default()

// 	// Use the logging middleware
// 	router.Use(LoggerMiddleware())

// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "pong"})
// 	})

// 	router.Run(":8080")
// }
