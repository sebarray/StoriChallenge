package transactionsRepository


import (
	 "storie/pkg/domain"

)
type Transaction interface {
	ReadTransaction( email string ) ([]domain.Mail, domain.Response)
}





var providers = map[string]Transaction{}

func init() {
	providers["query"] = Query{}
}

func GetProvider(provider string) Transaction {
	return providers[provider]
}
