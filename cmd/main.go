package main

import (
	"log"
	"net/http"

	"auction_api_gateway/proxy"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()


	authService := proxy.NewService("sign-up")
	proxy.RegisterServiceEndpoint(router, authService, "sign-up")
	proxy.RegisterGRPCServiceEndpoints(router, "/new-auction")

	
	server := &http.Server{
		Addr:    ":8050",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

