package handler

import (
	"storie/infrastructure/db/transactionsRepository"

	"storie/pkg/domain"
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
		return domain.Response{
			Message: "Error getting the data",
			Status:  500,
		}
	}
	CountTxn := len(txns)
	for i := 0; i < 3; i++ {
		txns = ServiceWokerTxns(txns)
		if len(txns) == 0 {
			break
		}
	}

	if CountTxn == len(txns) {
		return domain.Response{
			Message: "Error sending the email",
			Status:  500,
		}
	}

	var message string

	if reqBody.Emails == "" {
		message = "Email sent successfully to the user with their transaction history."
	} else {
		message = "Email sent successfully to all users with their corresponding transaction history."
	}
	if CountTxn > len(txns) {
		message = "Email sent successfully to the user with their transaction history. Error sending the email to some users."
		//implementar logica para  manejar las transacciones que no se pudieron enviar
	}

	return domain.Response{
		Message: message,
		Status:  200,
	}
}
