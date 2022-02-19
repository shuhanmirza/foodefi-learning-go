package main

import (
	"database/sql"
	"foodefi-go/fd-auth/api"
	db "foodefi-go/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/fd-db?sslmode=disable"
	address  = "0.0.0.0:9000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("can not start the server", err)
	}
}
