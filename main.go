package main

import (
	apis "aversan/BE/api"
	utilityFunc "aversan/BE/commonFunc"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	utilityFunc.StartApp();

	// INIZIALIZZAZIONE DB, FARE CAPO ALLE FUNZIONI DI UTILITA'
	db := utilityFunc.Dbinit()

	// API CHE GESTISCE IL NOSTRO MAIN
	// RITORNA TUTTI I DATI DI TUTTI GLI USER
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		apis.HandleUsers(w, r, db)
	})

	// API CHE RITORNA I DATI DI UN SOLO USER
	http.HandleFunc("/insertNewUser", func(w http.ResponseWriter, r *http.Request) {
		apis.InsertNewData(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
