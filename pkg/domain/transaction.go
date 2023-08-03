package domain


type Transaction struct {
    Id int `json:"id"`
	Date string `json:"date"`
	Amount float64 `json:"amount"`

}	

type Mail struct {
	SenderEmail string
	Transaction [] Transaction
	AverageDebit float64
	AverageCredit float64
	Link string
	Name string
	Balance float64
}



type Month struct {
Count int `json:"count"`
Name string `json:"Name"`
}