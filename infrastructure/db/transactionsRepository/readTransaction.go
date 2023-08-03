package transactionsRepository

import (
	"storie/pkg/domain"
	"storie/infrastructure/db"
  
   "log"
   "encoding/json"
)
type Query struct {
}


// ReadTransaction is a method to read the transaction of the user in the database
func (q Query) ReadTransaction( email string ) ([]domain.Mail, domain.Response) {
	mysl:=db.GetProvider("mysql")
	txn, err := mysl.TransactionDb()
	if err != nil {
		log.Println(err.Error())
		return nil, domain.Response{Status:500, Message: "transaction error"}
	}
	query:= "CALL get_txn(?)"	
	rows, err := txn.Query(query, email)
	if err != nil {
		log.Println(err.Error())
		return nil, domain.Response{Status:500, Message: "query error"}
	}
	defer rows.Close()
	var mails []domain.Mail
	for rows.Next() {
	var mail domain.Mail
	var txns string
	
		err := rows.Scan(&mail.SenderEmail,&txns,&mail.AverageDebit, &mail.AverageCredit ,&mail.Balance, &mail.Name)
		if err != nil {

			log.Println(err.Error())
			return nil, domain.Response{Status:500, Message: "scan error"}
		}

		err = json.Unmarshal([]byte(txns), &mail.Transaction)
	if err != nil {
		log.Println("Error al parsear el JSON:", err)
		return nil, domain.Response{Status:500, Message: "scan error"}
	}


	mails = append(mails, mail)
	}
return mails, domain.Response{Status:200, Message: "serch success"}
}



