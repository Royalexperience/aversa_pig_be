package main

import (
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"aversan/BE/commonFunc"
	"aversan/BE/api"
)
func main() {

	// INIZIALIZZAZIONE DB, FARE CAPO ALLE FUNZIONI DI UTILITA'
	db := utilityFunc.Dbinit();


	// API CHE GESTISCE IL NOSTRO MAIN
		// RITORNA TUTTI I DATI DI TUTTI GLI USER
		http.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request){
			getallUsersData.HandleUsers(w,r,db)
		})
	
		// API CHE RITORNA I DATI DI UN SOLO USER 
		//...	

	log.Fatal(http.ListenAndServe(":8080", nil))
}
