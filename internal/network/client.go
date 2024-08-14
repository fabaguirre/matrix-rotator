package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Matrix [][]int `json:"matrix"`
}

type Response struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
	Data    Data `json:"data"`
}

type Data struct {
	Max        int     `json:"max"`
	Min        int     `json:"min"`
	Average    float64 `json:"average"`
	TotalSum   int     `json:"totalSum"`
	IsDiagonal bool    `json:"isDiagonal"`
}

func GetMatrixStatistics(rotatedMatrix [][]int) (*Data, error) {
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

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	return &responseObject.Data, nil
}
