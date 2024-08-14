package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Matrix [][]int `json:"matrix"`
}

func GetMatrixStatistics(rotatedMatrix [][]int) (map[string]interface{}, error) {
	url := os.Getenv("API_URL")
	if url == "" {
		log.Fatalf("La variable de entorno NODE_API_URL no está definida")
	}

	requestBody := RequestBody{Matrix: rotatedMatrix}
	bodyData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url+"matrix/stats", "application/json", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error: received non-200 response status %d\n", response.StatusCode)
		return nil, err
	}

	var stats map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&stats); err != nil {
		return nil, err
	}

	return stats, nil
}
