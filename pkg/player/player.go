package player

import (
	"encoding/json"
	"errors"

	"github.com/Sourjaya/football-aws-lambda/pkg/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
)

var (
	ErrorFailedToUnmarshalRecord = "Failed to unmarshal record"
	ErrorFailedToFetchRecord     = "Failed to fetch record"
	ErrorInvalidPlayerData       = "Invalid Player data"
	ErrorInvalidID               = "Invalid ID"
	ErrorCouldNotMarshalItem     = "Could not marshal item"
	ErrorCouldNotDeleteItem      = "Could not delete item"
	ErrorCouldNotPostItem        = "Could not post item in DB"
	ErrorCouldNotPutItem         = "Could not put item in DB"
	ErrorPlayerAlreadyExists     = "Player already exists"
	ErrorPlayerDoesNotExist      = "Player does not exist"
	//ErrorGeneratingUUID          = "Could not generate UUID"
)

type Player struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Position  string `json:"position"`
	Club      string `json:"club"`
}

func GetPlayer(id, tableName string, client dynamodbiface.DynamoDBAPI) (*Player, error) {

	if !validators.IsValid(id) {
		return nil, errors.New(ErrorInvalidID)
	}

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := client.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(Player)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func GetPlayers(tableName string, client dynamodbiface.DynamoDBAPI) (*[]Player, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := client.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]Player)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func CreatePlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*Player,
	error,
) {
	var p Player

	if err := json.Unmarshal([]byte(req.Body), &p); err != nil {
		return nil, errors.New(ErrorInvalidPlayerData)
	}
	id := uuid.New()
	p.Id = id.String()

	currentPlayer, _ := GetPlayer(p.Id, tableName, client)
	if currentPlayer != nil && len(currentPlayer.Id) != 0 {
		return nil, errors.New(ErrorPlayerAlreadyExists)
	}

	item, err := dynamodbattribute.MarshalMap(p)

	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotPostItem)
	}
	return &p, nil
}

func UpdatePlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) (
	*Player,
	error,
) {

	var p Player
	if err := json.Unmarshal([]byte(req.Body), &p); err != nil {
		return nil, errors.New(ErrorInvalidID)
	}
	if !validators.IsValid(p.Id) {
		return nil, errors.New(ErrorInvalidID)
	}

	currentPlayer, _ := GetPlayer(p.Id, tableName, client)
	if currentPlayer != nil && len(currentPlayer.Id) == 0 {
		return nil, errors.New(ErrorPlayerDoesNotExist)
	}

	item, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotPutItem)
	}
	return &p, nil
}

func DeletePlayer(req events.APIGatewayProxyRequest, tableName string, client dynamodbiface.DynamoDBAPI) error {

	id := req.QueryStringParameters["id"]
	if !validators.IsValid(id) {
		return errors.New(ErrorInvalidID)
	}
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err := client.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
