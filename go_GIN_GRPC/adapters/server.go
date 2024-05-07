package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	pb "api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

var ErrInvalidValue = status.Error(codes.InvalidArgument, "invalid value type")

func (s *server) ProcessJSON(ctx context.Context, req *pb.Request) (*pb.Array, error) {
	// Unmarshal the JSON data from the request into a map of interfaces
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(req.JsonData), &jsonData); err != nil {
		return nil, err
	}

	// Transform the map[string]interface{} into the expected response format
	var responses []*pb.Response
	for key, value := range jsonData {
		// Convert the interface{} value to *structpb.Value
		val, err := structpb.NewValue(value)
		if err != nil {
			return nil, err
		}

		// Create a Response message for each key-value pair
		response := &pb.Response{
			ResultMap: map[string]*structpb.Value{
				key: val,
			},
		}
		responses = append(responses, response)
	}

	// Create and return the Array message containing the responses
	return &pb.Array{
		Msg: responses,
	}, nil
}

func main() {
	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterMyServiceServer(srv, &server{})

	log.Println("gRPC server is running on port :50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
