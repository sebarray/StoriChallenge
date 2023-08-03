package db



import (
	"database/sql"
	

	_ "github.com/go-sql-driver/mysql"
)





type DB interface {
	ConnectionDb() (*sql.DB, error)
	CheckDb() bool
	TransactionDb() (*sql.Tx, error)
}

var providers = map[string]DB{}

func init() {
	providers["mysql"] = Mysql{}
}

func GetProvider(provider string) DB {
	return providers[provider]
}
