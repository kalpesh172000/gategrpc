package main

import (
	"context"
	"log"
	"net"

	"github.com/kalpesh172000/gategrpc/services/common/gen/auth"
	"github.com/kalpesh172000/gategrpc/services/common/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type userServer struct {
	user.UnimplementedUserServiceServer
	authClient auth.AuthServiceClient
}

func (u *userServer) GetProfile(ctx context.Context, req *user.GetProfileRequest) (*user.GetProfileResponse, error) {
	/* res, err := u.authClient.Validate(ctx, &auth.ValidateRequest{Token: "valid-token"}) */
	res, err := u.authClient.Validate(ctx, &auth.ValidateRequest{Token: req.Token})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to validate token: %v", err)
	}

	if !res.Valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	return &user.GetProfileResponse{
		Profile: &user.Profile{
			Id:       req.Id,
			Name:     "john",
			Email:    "john@example.com",
			Password: "",
		},
	}, nil
}

func (s *userServer) CreateProfile(ctx context.Context, req *user.CreateProfileRequest) (*user.CreateProfileResponse, error) {
	// The `req` object is automatically populated with the data from the JSON body
	log.Printf("hello1 %v",req)
	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()
	return &user.CreateProfileResponse{
		Status: "Hello " + name + " " + email + " " + password,
	}, nil
}

func main() {
	conn, err := grpc.NewClient("localhost:10001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to attack userServer with authServer", err)
	}

	authClient := auth.NewAuthServiceClient(conn)

	// start grpc service
	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		log.Fatalln("failed to start the user service at 10002", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &userServer{authClient: authClient})
	log.Println("user service  started at 10002")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to start user's grpc Server", err)
	}
}
