package main

import (
	"log"
	"net/http"
	"net/http/httputil"

	"auction_api_gateway/proxy"

	"github.com/gin-gonic/gin"
)

func main() {
	auth_api_proxy := proxy.NewAuthApiProxy("localhost:8090")

	router := gin.New()

	router.POST("/sign-up", auth_api_proxy.ReverseAuthProxy("/auth/sign-up"))
	router.POST("/sign-in", auth_api_proxy.ReverseAuthProxy("/auth/sign-in"))

	server := &http.Server{
		Addr:    ":8050",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func ReverseProxy2() gin.HandlerFunc {
	target := "localhost:8090"

	return func(c *gin.Context) {
		c.Request.URL.Scheme = "http"
		c.Request.URL.Host = target
		c.Request.URL.Path = "/forbidden"
		director := func(req *http.Request) {
			req = c.Request
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
