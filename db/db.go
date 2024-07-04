package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaBd() *sql.DB {
	conexao := "user=postgres dbname=loja-golang password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
