package main

import (
	"os"

	l "github.com/Sourjaya/football-aws-lambda/logging"
	"github.com/Sourjaya/football-aws-lambda/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	client dynamodbiface.DynamoDBAPI
)

func main() {
	region := os.Getenv("AWS_REGION")
	Session, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	if err != nil {
		l.Error("Creating session")
		return
	}
	client = dynamodb.New(Session)
	lambda.Start(handler)
}

const tableName = "players"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetPlayer(req, tableName, client)
	case "POST":
		return handlers.CreatePlayer(req, tableName, client)
	case "PUT":
		return handlers.UpdatePlayer(req, tableName, client)
	case "DELETE":
		return handlers.DeletePlayer(req, tableName, client)
	default:
		return handlers.Unhandled()
	}
}
