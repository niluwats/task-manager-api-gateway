package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/niluwats/api-gateway/pkg/auth/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	authServiceUrl := fmt.Sprintf("localhost:%v", os.Getenv("AUTH_SERVICE"))
	conn, err := grpc.Dial(authServiceUrl, grpc.WithInsecure())

	if err != nil {
		log.Println("Couldn't connect : ", err)
	}

	defer conn.Close()

	return pb.NewAuthServiceClient(conn)
}
