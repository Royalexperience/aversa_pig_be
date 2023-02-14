package test

import (
	apis "aversan/BE/api"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	utilityFuncTest "aversan/BE/tests/utilityFuncForTesting"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

func TestInsertNewData(t *testing.T) {

    // Call a funzione di utilità per settare il db mockato
	//-------------------------------------------------------------
	db, mock, err := utilityFuncTest.SetUpMockDb(t)
	defer db.Close()
	//-------------------------------------------------------------


	// Configurazione del mock per l'inserimento di un nuovo utente (si aspetta che venga eseguita quella Query) e che ritorni risultato che abbia effetto
    // su una riga e una tabella
 	//-------------------------------------------------------------  
	mock.ExpectExec("INSERT INTO utenti").
		WithArgs("testuser", "testemail", "testpassword").
		WillReturnResult(sqlmock.NewResult(1, 1))
	//-------------------------------------------------------------


	// Creazione della richiesta HTTP con un body JSON valido
    //-------------------------------------------------------------
	requestBody := []byte(`{"username":"testuser","email":"testemail","password":"testpassword"}`)
	req, err := http.NewRequest("POST", "/insertNewUser", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Errore durante la creazione della richiesta HTTP: %s", err)
	}
	//-------------------------------------------------------------


	// Creazione del recorder per la risposta HTTP
    // da documentazione go ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests. 
	//-------------------------------------------------------------
	rr := httptest.NewRecorder()
	//-------------------------------------------------------------


	// Chiamata alla funzione InsertNewData passando la richiesta HTTP e il mock del database
	//-------------------------------------------------------------
	apis.InsertNewData(rr, req, db)
	//-------------------------------------------------------------


	// Verifica del risultato della chiamata
	//-------------------------------------------------------------
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Errore nella richiesta HTTP: il server ha restituito lo status code %v invece di %v", status, http.StatusCreated)}
    //-------------------------------------------------------------

}
