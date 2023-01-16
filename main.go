package main

import (
	"database/sql"
	"fmt"
	"log"
	"simplebank/api"
	"simplebank/db/store"
	"simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err1 := util.LoadConfig(".")
	fmt.Println(config)
	fmt.Println(err1)
	if err1 != nil {
		log.Fatal("cant load the config file due to :", err1.Error())

	}
	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable")
	if err != nil {
		fmt.Println("got error while opening the database connection ", err.Error())
		//panic(err)

	}

	newServerConnection := api.NewServer(store.NewStore(db))

	err = newServerConnection.Startserver(config.SAddress)
	if err != nil {
		log.Fatal("can't connect to the db")
		panic(err)
	}

}
