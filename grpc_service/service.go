package grpc_service

import (
	"auction_api_gateway/grpc_service/models"
	"auction_api_gateway/grpc_service/pb"
	"context"
	"log"

	"google.golang.org/grpc"
)




func NewAuctionRequest(reqData *models.NewAuctionRequestModel, host string) (*models.NewAuctionResponseModel, error){
	opts := []grpc.DialOption{
        grpc.WithInsecure(),
    }
	conn, err := grpc.Dial(host, opts...) //example: 127.0.0.1:5300
	if err != nil{
		log.Println(err)
		return nil, err
	}

	defer conn.Close()

	client := pb.NewAuctionServiceClient(conn)

	
	req := &pb.NewAuctionRequest{
		UserId: reqData.UserId,
		Title: reqData.Title,
		Type: reqData.Type,
		Status: reqData.Status,
		Date: reqData.Date,
	}

	res, err := client.NewAuction(context.Background(), req)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	return &models.NewAuctionResponseModel{
		AuctionId: res.AuctionId,
	}, nil
	
}