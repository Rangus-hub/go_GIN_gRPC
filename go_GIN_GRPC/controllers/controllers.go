package main

import (
	pb "api/proto" // Import your gRPC package
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Define a struct to hold your gRPC client connection
type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.MyServiceClient
}

func main() {
	// Initialize your gRPC client connection
	grpcClient := initGRPCClient()

	// Initialize Gin router
	router := gin.Default()

	// Define Gin routes
	router.POST("/process", func(c *gin.Context) {
		// Parse JSON data from the request
		var jsonData map[string]interface{}
		if err := c.BindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Transform JSON data into the format expected by gRPC server
		req := &pb.Request{
			JsonData: jsonData["json_data"].(string),
		}

		// Send request to gRPC server
		response, err := grpcClient.client.ProcessJSON(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Transform response into the format expected by HTTP clients
		responseData := make(map[string]interface{})
		for _, resp := range response.Msg {
			for key, value := range resp.ResultMap {
				// You can handle each field in the response here
				responseData[key] = value.String()
			}
		}

		// Send response back to HTTP client
		c.JSON(http.StatusOK, responseData)
	})

	// Run the Gin server
	router.Run(":8080")
}

// Function to initialize gRPC client connection
func initGRPCClient() *GRPCClient {
	// Dial gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}

	// Create gRPC client
	client := pb.NewMyServiceClient(conn)

	// Return gRPC client instance
	return &GRPCClient{
		conn:   conn,
		client: client,
	}
}

// type Data struct {
// 	Id   string `json : "id"`
// 	Name string `json : "name"`
// }

// type Controller struct {
// 	repo repository.Repository
// }

// func NewController(repo repository.Repository) *Controller {
// 	return &Controller{repo: repo}
// }

// func PostToDB() string {
// 	return "this is data from db"
// }
// func (d *Controller) Deletedata(c *gin.Context) {
// 	name := c.Query("name")
// 	id := c.Query("id")
// 	s := d.repo.Deletedata(name, id)
// 	c.JSON(200, gin.H{
// 		"data": s,
// 	})
// }
// func (d *Controller) AddData(c *gin.Context) {
// 	var requestData Data
// 	if err := c.BindJSON(&requestData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	print("data " + requestData.Id)
// 	s := d.repo.AddData(requestData.Name, requestData.Id)
// 	c.JSON(200, gin.H{
// 		"data": s,
// 	})
// }
