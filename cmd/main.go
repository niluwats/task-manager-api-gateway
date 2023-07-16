package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth"
)

func main() {
	r := gin.Default()
	auth.RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%v", os.Getenv("WEB_PORT")))
}
