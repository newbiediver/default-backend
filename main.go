package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	quitSignal := make(chan os.Signal)

	initWebService()

	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
}

func initWebService() {
	server := gin.New()
	server.Use(customLogger(), gin.Recovery())

	port := ":8000"

	server.StaticFile("/", "web/404.html")
	//server.StaticFile("/css/style.css", "web/css/style.css")
	go func() {
		fmt.Printf("[System] Starting web service. port%s", port)
		if err := http.ListenAndServe(port, server); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func customLogger() gin.HandlerFunc {
	out := gin.DefaultWriter

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()

		latency := end.Sub(start)
		proxyHeader := c.GetHeader("X-Forwarded-For")

		var clientAddress string

		if proxyHeader != "" {
			ips := strings.Split(proxyHeader, ",")
			if len(ips) > 0 {
				clientAddress = ips[0]
			} else {
				clientAddress = proxyHeader
			}
		} else {
			clientAddress = c.ClientIP()
		}

		timeString := time.Now().In(time.UTC).Format("2006/01/02 - 15:04:05")
		format := fmt.Sprintf("[System] %v | %3d | %13v | %15s | %-7s %#v\n",
			timeString,
			c.Writer.Status(),
			latency,
			clientAddress,
			c.Request.Method,
			c.Request.RequestURI,
		)

		_, _ = fmt.Fprintf(out, format)
	}
}
