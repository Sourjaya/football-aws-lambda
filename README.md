# GoLang Player API with DynamoDB Integration

This is a simple REST API written in GoLang that provides CRUD (Create, Read, Update, Delete) operations for managing player information. The data is stored in Amazon DynamoDB, and the API is deployed as an AWS Lambda function, accessible through the AWS API Gateway.

## Features

- **Get All Players**: Retrieve a list of all players stored in the DynamoDB table.
- **Get Player by ID**: Retrieve player information by providing the player's ID as a URL parameter.
- **Create Player**: Add a new player to the DynamoDB table by sending player information as a JSON object in the request body.
- **Delete Player by ID**: Delete a player from the DynamoDB table using the player's ID provided as a URL parameter.
- **Update Player by ID**: Update existing player information in the DynamoDB table using the player's ID provided as a URL parameter.

## Data Structure

Each player record in the DynamoDB table consists of the following fields:

- **id**: Unique identifier for the player.
- **firstName**: First name of the player.
- **lastName**: Last name of the player.
- **country**: Country of the player.
- **club**: Current club or team of the player.

## Usage
To use this API, you can interact with it by sending ID or data as URL parameters using HTTP requests. Below are the available endpoints:

| HTTP Method | EndPoint         | Description              |
|-------------|------------------|--------------------------|
| GET         | /staging         | Retrieve all players.    |
| GET         | /staging?id={ID} | Retrieve a player by ID. |
| POST        | /staging         | Create a new player.     |
| PUT         | /staging?id={ID} | Update a player by ID.   |
| DELETE      | /staging?id={ID} | Delete a player by ID.   |
## Example Request

```
GET /staging?id=b68c99b8-c770-412e-aff6-36aa977b563
```
## Example Response
```json
{
    "id":"b68c99b8-c770-412e-aff6-36aa977b563",
    "firstName": "Lionel",
    "lastName": "Messi",
    "country": "arg",
    "position": "FW",
    "club": "Inter Miami"
}
```
## Deployment
### 1. Build the executable
```bash
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go
```
### 2. Make a zip
```bash
zip -jrm myFunction.zip bootstrap
```
### 3. Upload to AWS lambda
The API is deployed as an AWS Lambda function, integrated with the AWS API Gateway for access. 
> **NOTE**
> Ensure that appropriate permissions(execution roles)are set for the Lambda function to interact with DynamoDB.

Development
To set up this project locally for development or testing purposes, follow these steps:

1. Clone the repository.
2. Install GoLang and required dependencies.
3. Set up your AWS credentials and configure the AWS SDK v1 for GoLang.
4. Customize the code as needed.
5. Deploy the Lambda function and API Gateway using the AWS CLI or preferred deployment method.

The complete guide can be found in this [article](https://dev.to/sourjaya/build-and-deploy-rest-api-with-aws-lambda-api-gateway-and-dynamodb-using-golang-58ap).
