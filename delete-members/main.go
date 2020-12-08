package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request)
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Hello, Go lambda"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
