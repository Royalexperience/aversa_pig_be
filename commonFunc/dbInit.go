package utilityFunc

import(
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"os"
	"github.com/joho/godotenv"
)
// CARICO DALLE VARIABILI DI AMBIENTE I VALORI, LI INSERISCO NELLA STRINGA DI CONNESSIONE E MI CONNETTO, IN CASO DI ERRORE AMEN ALTRIMENTI RITORNO IL PUNTATORE 
func Dbinit () *sql.DB {

	//VARIABILI D'AMBIENTE//
	godotenv.Load("resources/config-local.env")
	USERSDB_USERNAME:= os.Getenv("USERSDB_USERNAME")
	USERSDB_PASSWORD:= os.Getenv("USERSDB_PASSWORD")
	USERSDB_IP:= os.Getenv("USERSDB_IP")
	USERSDB_NAME := os.Getenv("USERSDB_NAME")
	//VARIABILI D'AMBIENTE//

	// STRINGA DI CONNESSIONE 
	CONNECTION_STRING := "postgres"+"://"+USERSDB_USERNAME+":"+USERSDB_PASSWORD+"@"+USERSDB_IP+"/"+USERSDB_NAME+"?sslmode=disable"

	// MI CONNETTO 
	db, err := sql.Open("postgres", CONNECTION_STRING)
	
	// CASO DI ERRORE
	if err != nil {
		log.Fatal(err)
	}
	
	return db
	
}
