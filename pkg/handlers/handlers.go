// This code is part of handlers package
package handlers

// import other packages
import (
	"net/http"

	l "github.com/Sourjaya/football-aws-lambda/logger"
	"github.com/Sourjaya/football-aws-lambda/pkg/player"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// create a error message variable
var ErrorMethodNotAllowed = "method not allowed"

// ErrorBody represents the structure of the error response body.
type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

// this function will be called when the method type is "GET"
func GetPlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	// if the id is provided as URL parameter then call GetPlayerByName function from player package.
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		l.Info("ID given")
		result, err := player.GetPlayer(id, tableName, client)
		if err != nil {
			return response(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return response(http.StatusOK, result)
	}
	// else call GetPlayers function from player package.
	result, err := player.GetPlayers(tableName, client)
	if err != nil {
		return response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return response(http.StatusOK, result)

}

// this function will be called when the method type is "POST"
func CreatePlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	// Call CreatePlayer function from player package.
	result, err := player.CreatePlayer(req, tableName, client)
	if err != nil {
		return response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return response(http.StatusCreated, result)
}

// this function will be called when the method type is "PUT"
func UpdatePlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := player.UpdatePlayer(req, tableName, client)
	if err != nil {
		return response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return response(http.StatusOK, result)
}

// this function will be called when the method type is "DELETE"
func DeletePlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	err := player.DeletePlayer(req, tableName, client)

	if err != nil {
		return response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return response(http.StatusOK, nil)
}

// Function that return a Method not allowed error message for those unhandled method types.
func Unhandled() (*events.APIGatewayProxyResponse, error) {
	return response(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
