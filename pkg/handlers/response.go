package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func response(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Content-Type", "Access-Control-Allow-Methods": "OPTIONS, POST, GET, PUT, DELETE", "Access-Control-Allow-Credentials": "true"}}
	resp.StatusCode = status

	sBody, _ := json.Marshal(body)
	resp.Body = string(sBody)
	return &resp, nil
}
