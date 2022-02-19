package main

import (
	"database/sql"
	"fmt"
	"foodefi-go/fd-auth"
	db "foodefi-go/sqlc"
	"foodefi-go/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	globalConfig, err := util.LoadGlobalConfig(".")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	fmt.Println(globalConfig)

	conn, err := sql.Open(globalConfig.DBDriver, globalConfig.DBSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	store := db.NewStore(conn)

	fdAuthServer := fd_auth.NewServer(store)
	err = fdAuthServer.Start(globalConfig.FdAuthServerAddress)
	if err != nil {
		log.Fatal("can not start the fdAuthServer", err)
	}
}
