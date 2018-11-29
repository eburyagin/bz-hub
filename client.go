package main

import (
	"bz-client/api"
	"bz-lib/app"
	"encoding/json"
	"log"
	"net/http"
)

func listClients(w http.ResponseWriter, r *http.Request, app *app.AppContext) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	resp, err := api.ListClients_v1(app)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(resp.Data)
}
