package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	types "github.com/FluxNFTLabs/sdk-go/chain/indexer/campclash"
)

func main() {
	cc, err := grpc.Dial("localhost:4464", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer cc.Close()
	if err != nil {
		panic(err)
	}
	client := types.NewCampclashQueryClient(cc)

	stream, err := client.SubscribeUserActivity(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}
