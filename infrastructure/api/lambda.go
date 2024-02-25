package api

import (
	"log"
	// "context"
	// "encoding/json"
	// "fmt"

	"storie/infrastructure/api/controller"

	"github.com/aws/aws-lambda-go/lambda"
)

func Start() {
	controller := controller.GetProvider()

	log.Println("Starting application")
	lambda.Start(controller.SendMail)
}
