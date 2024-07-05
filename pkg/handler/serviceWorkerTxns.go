package handler

import (
	"log"
	"storie/infrastructure/db/storageFile"
	"storie/infrastructure/service"
	"storie/pkg/domain"
	"sync"
)

func ServiceWokerTxns(txns []domain.Mail) []domain.Mail {
	numWorkers := 10
	tasks := make(chan domain.Mail, len(txns))
	errors := make(chan domain.Mail, len(txns)) // Canal para errores de transacción

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, errors, &wg)
	}

	for _, txn := range txns {
		tasks <- txn
	}
	close(tasks)

	wg.Wait()
	close(errors)

	return handleErrors(errors)
}

func worker(tasks <-chan domain.Mail, errors chan<- domain.Mail, wg *sync.WaitGroup) {
	defer wg.Done()
	for txn := range tasks {
		if err := processTransaction(txn); err != nil {
			errors <- txn // Enviar la transacción y el error al canal de errores
		}
	}
}

func processTransaction(txn domain.Mail) error {
	gmail := service.GetProvider("gmail")
	storage := storageFile.GetProvider("s3")

	link, err := storage.UploadTransactions(txn.Transaction, txn.SenderEmail)
	if err != nil {
		return err // Retornar el error
	}
	txn.Link = link

	if err := gmail.Send(txn); err != nil {
		return err
	}

	return nil
}

func handleErrors(errors <-chan domain.Mail) []domain.Mail {
	var txnsErr []domain.Mail
	for txnErr := range errors {
		txnsErr = append(txnsErr, txnErr)
		log.Printf("Error processing transaction for %s: %v\n", txnErr.Name, txnErr.Err)
	}
	return txnsErr
}
