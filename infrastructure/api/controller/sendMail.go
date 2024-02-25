package controller

import (
	"encoding/json"
	"log"
	"storie/pkg/domain"
	"storie/pkg/handler"

	"github.com/aws/aws-lambda-go/events"
)

type SendMailController struct {
}

type ISendMailController interface {
	SendMail(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func GetProvider() ISendMailController {
	return SendMailController{}
}

func (S SendMailController) SendMail(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody domain.Request
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		log.Println("Error parsing request body:", err.Error())
		return events.APIGatewayProxyResponse{
			Body:       "bad request",
			StatusCode: 400,
		}, nil
	}
	handler := handler.GetProvider()

	response := handler.SendMail(reqBody)

	return events.APIGatewayProxyResponse{
		Body:       response.Message,
		StatusCode: response.Status,
	}, nil

}
