package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kalpesh172000/gategrpc/services/common/gen/user"
	"github.com/kalpesh172000/gategrpc/services/common/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := runtime.NewServeMux(runtime.WithErrorHandler(util.CustomHttpError),)

	//connect to user service it is  REST exposed

	err := user.RegisterUserServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:10002",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)

	if err != nil {
		log.Fatalln("failed to connect grpc gateway to grpc Service",err)
	}

	log.Println("Gateway running one :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("failed to start grpc-gateway",err)
	}
}
