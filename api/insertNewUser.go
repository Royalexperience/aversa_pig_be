package apis

import (
	"database/sql"
	"encoding/json"
	"net/http"
	userData "aversan/BE/types"
	utilsConstants "aversan/BE/utils"
	"github.com/go-playground/validator/v10"
	"aversan/BE/commonFunc"
)

// FUNZIONE PRINCIPALE PER INSERIMENTO DELL' UTENTE 
func InsertNewData(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// CHECK SUL METODO
	//-------------------------------------------------------------
	 if(utilityFunc.CheckMethod(r,w,utilsConstants.METHOD_POST)){return}
	//-------------------------------------------------------------

	// DECODIFICA IL JSON ARRIVATO NELLA STRUCT
	//-------------------------------------------------------------
	var user userData.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	//-------------------------------------------------------------

	// CALL A FUNZIONi PER VALIDARNE I CAMPI (VIENE UTILIZZATA LIBRERIA + FUNZIONI CUSTOM)
	//-------------------------------------------------------------
	validate := validator.New()
	if validate.Struct(user) != nil {

		handleResponse(w,findErrorMessage(user))
		return
	}
	//-------------------------------------------------------------


	// Assegno la query che devo fare + prendo l'eventuale errore dalla query il "_,"  serve perché quella funz db.Exec ritorna due valori ma a me serve solo uno 
		//dunque per non mandare in errore il compilatore poiché va in erro per variabili non usate si usa quel "_," 
	//-------------------------------------------------------------
	sqlStatement := utilsConstants.QUERY_INSERT_NEW_DATA
	_, err = db.Exec(sqlStatement, user.Username, user.Email, user.Password)
	if err != nil {
		utilityFunc.QueryErrorResponse(err,w ,err.Error())
		return
	}
	//-------------------------------------------------------------

	
	// Creo la risposta tramite una mappa e la mando in risposta per eventuale 201ù
	//-------------------------------------------------------------
	response := map[string]string{"message": "User created successfully"}
	responseBytes, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBytes)
	//-------------------------------------------------------------

}


func handleResponse(w http.ResponseWriter, message string) {
	// Creo mappa per eventuale bad request e la mando in risposta (utilizzata funzione custom ErrorParsingResponse per far felice il compilatore)
	//-------------------------------------------------------------
	response := map[string]string{
		"error":   utilsConstants.ERROR_400_BAD_REQUEST,
		"message": message,
	}
	responseBytes, errorResponse := json.Marshal(response)
	utilityFunc.ErrorParsingResponse(errorResponse,w);
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(responseBytes)
	//-------------------------------------------------------------
}

func findErrorMessage(user userData.User) string {

	// Check su ciascuno dei campi se vuoti oppure body sbagliata e ritorno stringa errore
	//-------------------------------------------------------------
   	var errorMessage string ;
	if (user.Email == "") {
		errorMessage = utilsConstants.EMAIL_FIELD_MISSING
	}
	if (user.Username == "") {
		errorMessage = utilsConstants.USERNAME_FIELD_MISSING
	}
	if (user.Password == "") {
		errorMessage = utilsConstants.PASSWORD_FIELD_MISSING
	}

	if (user.Email == "" && user.Username == "" && user.Password == "" ){
		errorMessage = utilsConstants.ERROR_MESSAGE_EMPTY_BODY_OR_WRONG_FIELDS
	}
	return errorMessage
	//-------------------------------------------------------------

}
