package handler

import (
	"log"
	"storie/infrastructure/db/storageFile"
	"storie/infrastructure/db/transactionsRepository"

	"encoding/json"

	"storie/infrastructure/service"
	"storie/pkg/domain"
	"sync"

	"github.com/aws/aws-lambda-go/events"
)

func SendMail(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody domain.Request
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err.Error())
		return events.APIGatewayProxyResponse{
			Body:       "Error al analizar el cuerpo de la solicitud",
			StatusCode: 500,
		}, nil
	}

	repoTxn := transactionsRepository.GetProvider("query")

	txns, Response := repoTxn.ReadTransaction(reqBody.Emails)

	if Response.Status != 200 {
		return events.APIGatewayProxyResponse{
			Body:       "Error al obtener los datos",
			StatusCode: 500,
		}, nil
	}
	gmail := service.GetProvider("gmail")
	storage := storageFile.GetProvider("s3")

	var wg sync.WaitGroup

	for _, txn := range txns {

		wg.Add(1)

		go func(txn domain.Mail) {
			defer wg.Done()

			var err error
			txn.Link, err = storage.UploadTransactions(txn.Transaction, txn.SenderEmail)
			if err != nil {
				log.Println(err.Error())

				return
			}
			err = gmail.Send(txn)
			if err != nil {
				log.Println(err.Error())

				return
			}
		}(txn)
	}

	wg.Wait()

	return events.APIGatewayProxyResponse{
		Body:       "Correo enviado correctamente",
		StatusCode: 200,
	}, nil

}
