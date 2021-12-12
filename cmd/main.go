package main

import (
	"log"
	"net/http"

	"auction_api_gateway/proxy"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()


	signUp := proxy.NewService("sign-up")
	proxy.RegisterServiceEndpoint(router, signUp, "/sign-up")

	signIn := proxy.NewService("sign-in")
	proxy.RegisterServiceEndpoint(router, signIn, "/sign-in")

	

	api := router.Group("api")
	api.Use(proxy.ServeMiddleware)
	{
		proxy.RegisterGRPCServiceEndpoints(api, "/new-auction")
	}

	
	server := &http.Server{
		Addr:    ":8050",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

