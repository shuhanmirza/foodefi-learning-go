package main

import (
	"database/sql"
	"fmt"
	fd_auth "foodefi-go/fd-auth"
	fd_event_listener "foodefi-go/fd-event-listener"
	db "foodefi-go/sqlc"
	"foodefi-go/util"
	_ "github.com/lib/pq"
	"log"
	"sync"
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

	// TODO: restructure these servers to run as separate microservices

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		fdAuthServer := fd_auth.NewServer(store)
		err = fdAuthServer.Start(globalConfig.FdAuthServerAddress)
		if err != nil {
			log.Fatal("can not start the fdAuthServer", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fdEventListener := fd_event_listener.NewServer(store)
		err = fdEventListener.Start(globalConfig.FdEventListenerAddress)
		if err != nil {
			log.Fatal("can not start the fdEventListener", err)
		}
	}()

	wg.Wait()

}
