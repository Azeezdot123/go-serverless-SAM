package main

import (
	"github.com/azeezdot123/go-serverless-sam/handlers"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

const tableName = "todo"

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	if err != nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetTodo(req, tableName, dynaClient)

	// case "POST":
	// 	return handlers.CreateTodo(req, tableName, dynaClient)

	// case "PUT":
	// 	return handlers.UpdateTodo(req, tableName, dynaClient)

	// case "DELETE":
	// 	return handlers.DeleteTodo(req, tableName, dynaClient)

	default:
		return handlers.UnhandledMethod()
	}
}
