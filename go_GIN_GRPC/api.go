package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// type intr interface {
// }

func main() {

	//jsonData := `{"user": 1, "name": "Raju"}`

	// Prepare JSON data
	jsonData := map[string]interface{}{
		"json_data": `{"data":"DATA TADA"}`,
	}
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	// Send HTTP POST request to the server
	resp, err := http.Post("http://localhost:8080/process", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	// Print response
	fmt.Println("Response:", response)
}
