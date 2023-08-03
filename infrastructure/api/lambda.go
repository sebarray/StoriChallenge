package api

import (
	"log"
	// "context"
	// "encoding/json"
	// "fmt"

	"storie/pkg/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func Start() {

	log.Println("Starting application")
	lambda.Start(handler.SendMail)
}
