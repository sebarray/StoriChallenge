package handler

import (
	"log"
	"storie/infrastructure/db/storageFile"
	"storie/infrastructure/db/transactionsRepository"

	"storie/infrastructure/service"
	"storie/pkg/domain"
	"sync"
)

type IHandler interface {
	SendMail(reqBody domain.Request) domain.Response
}

type Handler struct {
}

func GetProvider() IHandler {
	return &Handler{}
}

func (h *Handler) SendMail(reqBody domain.Request) domain.Response {

	repoTxn := transactionsRepository.GetProvider("query")

	txns, Response := repoTxn.ReadTransaction(reqBody.Emails)

	if Response.Status != 200 {
		Response.Message = "Error getting the data"
		return Response
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

	if reqBody.Emails == "" {
		Response.Message = "Email sent successfully to all users with their corresponding transaction history."
	} else {
		Response.Message = "Email sent successfully to the user with their transaction history."
	}

	return Response
}
