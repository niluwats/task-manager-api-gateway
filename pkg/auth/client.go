package auth

import (
	"log"

	"github.com/niluwats/api-gateway/pkg/auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	// authServiceUrl := fmt.Sprintf("dns:///auth-service:%v", os.Getenv("AUTH_SVC_PORT"))
	authServiceUrl := "auth-service:50051"
	conn, err := grpc.Dial(authServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Println(conn)
	log.Println(conn.GetState())
	if err != nil {
		log.Println("Couldn't connect : ", err)
	}

	// defer conn.Close()
	return pb.NewAuthServiceClient(conn)
}
