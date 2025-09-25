package main 

import (
	"context"
	"log"
	"net"

	"github.com/kalpesh172000/gategrpc/services/common/gen/auth"
	"google.golang.org/grpc"
)

type authServer struct {
	auth.UnimplementedAuthServiceServer
}

func (a *authServer) Validate (ctx context.Context, req *auth.ValidateRequest)(*auth.ValidateResponse, error){
	return &auth.ValidateResponse{Valid: req.Token == "valid-token"},nil 
}


func main(){
	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatalln("failed to start auth serves at :10001",err)
	}

	grpcServer := grpc.NewServer()


	auth.RegisterAuthServiceServer(grpcServer, &authServer{})
	log.Println("started grpc server on 10001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to register the AuthServiceServer", err)
	}
}
