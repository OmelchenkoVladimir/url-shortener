package main

import (
	"log"
	"url_shortener_main/config"
	"url_shortener_main/db"
	"url_shortener_main/routing"

	"github.com/valyala/fasthttp"
)

func main() {
	conf, err := config.ReadConfig("./config/config.json")
	if err != nil {
		log.Fatal("Can't read configuration file\n", err.Error())
	}
	dbpath := conf.Db.Host + ":" + conf.Db.Port
	database, err := db.NewDB(dbpath)
	if err != nil {
		log.Fatal("Can't establish database connection\n", err.Error())
	}
	defer database.Client.Close()
	address := conf.Server.Url + ":" + conf.Server.Port
	router := routing.New(address, database)
	log.Fatal(fasthttp.ListenAndServe(address, router.Handler))
	log.Println("END")
}
