package task

import (
	"log"

	"github.com/niluwats/api-gateway/pkg/task/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.TaskServiceClient
}

func InitServiceClient() pb.TaskServiceClient {
	taskServiceUrl := "task-service:50052"
	conn, err := grpc.Dial(taskServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("Couldn't connect to task service : ", err)
	}

	defer conn.Close()

	return pb.NewTaskServiceClient(conn)
}
