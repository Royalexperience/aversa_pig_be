package test

import (
	utilityFuncTest "aversan/BE/tests/utilityFuncForTesting"
	"net/http"
	"net/http/httptest"
	"testing"
	apis "aversan/BE/api"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllUsers(t *testing.T) {

	// Call a funzione di utilità per settare il db mockato
	//-------------------------------------------------------------
	db, mock, err := utilityFuncTest.SetUpMockDb(t)
	defer db.Close()
	//-------------------------------------------------------------


	// Inserimento dati fittizzi all'interno del mock db
	//-------------------------------------------------------------
	rows := sqlmock.NewRows([]string{"email", "username", "password"}).
	AddRow("user1@example.com", "user1", "password1").
	AddRow("user2@example.com", "user2", "password2")
	//-------------------------------------------------------------


	// Cosa ci si aspetta
	//-------------------------------------------------------------
	mock.ExpectQuery("SELECT email, username, password FROM utenti").WillReturnRows(rows)
	//-------------------------------------------------------------

	// Creazione richiesta 
	//-------------------------------------------------------------
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Errore durante la creazione della richiesta HTTP: %s", err)
	}
	//-------------------------------------------------------------


	// Creazione del recorder per la risposta HTTP
	// da documentazione go ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.
	//-------------------------------------------------------------
	rr := httptest.NewRecorder()
	//-------------------------------------------------------------


	// Chiamata alla funzione HandleUsers passando la richiesta HTTP e il mock del database
	//-------------------------------------------------------------
	apis.HandleUsers(rr, req, db)
	//-------------------------------------------------------------


	// Verifica del risultato della chiamata
	//-------------------------------------------------------------
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Errore nella richiesta HTTP: il server ha restituito lo status code %v invece di %v", status, http.StatusOK)}
	//-------------------------------------------------------------
}
