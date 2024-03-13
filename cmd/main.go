// main package
package main

// import other packages
import (
	"os"

	l "github.com/Sourjaya/football-aws-lambda/logger"
	"github.com/Sourjaya/football-aws-lambda/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Declare a variable of interface type dynamodbiface.
// DynamoDBAPI that enables to mock the DynamoDB service client.
var (
	client dynamodbiface.DynamoDBAPI
)

func main() {
	// get the aws region from the environment variable
	region := os.Getenv("AWS_REGION")

	// Create a new Session passing region as configuration
	Session, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	if err != nil {
		l.Error("Creating session")
		return
	}

	// Create a new instance of DynamoDB client
	client = dynamodb.New(Session)

	// Start takes a handler and talks to an internal Lambda endpoint to pass requests to the handler.
	lambda.Start(handler)
}

// DynamoDB table
const tableName = "players"

// Handler function
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		// if HTTP method type is GET
		return handlers.GetPlayer(req, tableName, client)
	case "POST":
		// if HTTP method type is POST
		return handlers.CreatePlayer(req, tableName, client)
	case "PUT":
		// if HTTP method type is PUT
		return handlers.UpdatePlayer(req, tableName, client)
	case "DELETE":
		// if HTTP method type is DELETE
		return handlers.DeletePlayer(req, tableName, client)
	default:
		// for all other HTTP method types
		return handlers.Unhandled()
	}
}
