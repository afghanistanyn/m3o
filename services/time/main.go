package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/m3o/services/time/handler"
	pb "github.com/m3o/m3o/services/time/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("time"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTimeHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
