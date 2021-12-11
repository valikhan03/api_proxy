package proxy

import (
	"auction_api_gateway/grpc_service"
	"auction_api_gateway/grpc_service/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuction(c *gin.Context){
	var input models.NewAuctionRequestModel
	c.BindJSON(&input)

	res, err := grpc_service.NewAuctionRequest(&input, "127.0.0.1:50051")
	if err != nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&res)
	if err != nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, data)
}


func RegisterGRPCServiceEndpoints(router *gin.Engine, path string){
	router.POST(path, NewAuction)
}





