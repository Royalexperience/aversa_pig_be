package utilityFunc

import (
	utilsConstants "aversan/BE/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func QueryErrorResponse(errorResponse error, w http.ResponseWriter, queryErrorMessage string) {
	var errorString string = utilsConstants.ERROR_500
	var statusCode int = http.StatusInternalServerError
	// SE IL MESSAGGIO DI ERRORE CONTIENE LA SOTTOSTRINGA "duplicate key"
	if strings.Contains(queryErrorMessage, utilsConstants.DUPLICATE_KEY_STRING) {
		// IMPOSTA LO STATUS CODE COME 409
		statusCode = http.StatusConflict;
		// IMPOSTA L'ERRORE COME 409
		errorString = utilsConstants.ERROR_409
		// IMPOSTA IL MESSAGGIO DI ERRORE CHE COMUNICA CHE L'EMAIL E' GIA' REGISTRATA
		queryErrorMessage = utilsConstants.ERROR_409_EMAIL_ALREADY_REGISTERED
	}
	// COSTRUISCI LA MAPPA CHE RAPPRESENTERA' IL JSON DI RISPOSTA
	response := map[string]string{
		"error":   errorString,
		"message": queryErrorMessage,
	}
	// CODIFICA LA MAPPA IN FORMATO JSON
	responseBytes, err := json.Marshal(response)
	// VERIFICA SE CI SONO ERRORI DERIVANTI DALLA CODIFICA DEL JSON
	ErrorParsingResponse(err, w)
	// IMPOSTA L'HEADER, SPECIFICANDO CHE IL MESSAGGIO DI RISPOSTA SARA'
	// IN FORMATO JSON
	w.Header().Set("Content-Type", "application/json")
	// INVIA LO STATUS CODE
	w.WriteHeader(statusCode)
	// INVIA IL MESSAGGIO DI RISPOSTA
	w.Write(responseBytes)
}
