package main

import (
	"bz-lib/app"
	"bz-lib/cfg"
	"flag"
	"log"
	"net/http"

	nats "github.com/nats-io/go-nats"
)

var cr = flag.String("c", "./config.json", "конфигурационный файл")

func main() {

	app := new(app.AppContext)

	flag.Parse()

	var err error
	app.Conf, err = cfg.Load(*cr)
	if err != nil {
		log.Fatalln("Ошибка загрузки конфигурации!")
	}

	c, err := nats.Connect(app.Conf.Bus.Urls)
	if err != nil {
		log.Fatal(err)
	}
	app.Nats, err = nats.NewEncodedConn(c, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		listClients(w, r, app)
	})
	http.ListenAndServe(":"+app.Conf.Http.Port, nil)
}
