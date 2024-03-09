package handlers

import (
	"net/http"

	l "github.com/Sourjaya/football-aws-lambda/logging"
	"github.com/Sourjaya/football-aws-lambda/pkg/player"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetPlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {

	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		l.Info("ID given")
		result, err := player.GetPlayer(id, tableName, client)
		if err != nil {
			return response(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return response(http.StatusOK, result)
	}
	result, err := player.GetPlayers(tableName, client)
	if err != nil {
		return response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return response(http.StatusOK, result)

}

func CreatePlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := player.CreatePlayer(req, tableName, client)
	if err != nil {
		return response(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return response(http.StatusCreated, result)
}

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

func Unhandled() (*events.APIGatewayProxyResponse, error) {
	return response(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
