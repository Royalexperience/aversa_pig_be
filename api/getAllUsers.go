package apis

import "net/http"
import "encoding/json"
import "database/sql"
import "aversan/BE/types"
import "aversan/BE/utils"
import "aversan/BE/commonFunc"

// FUNZIONE CHE RESTITUISCE TUTTI I DATI DEGLI USER
func HandleUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	
	// GESTISCE ERRORE SE IL METODO RICHIAMATO PER QUELL' API NON E' UNA GET
	if(utilityFunc.CheckMethod(r,w,utilsConstants.METHOD_GET)){
		return
	 }
	
	// ESEGUE LA QUERY E IN CASO DI ERRORE RESTITUISCE 500
	rows, err := db.Query(utilsConstants.QUERY_ALL_DATA)
	if err != nil {
		utilityFunc.QueryErrorResponse(err,w ,err.Error())
		return
	}
    // CHIUDE CLOSE PER EVITARE PROBLEMI CON IL DATABASE
	defer rows.Close()

	// ASSOCIA AD USERS UN ARRAY DELLA STRUCT DICHIARATA NELL' ALTRA CARTELLA
	users := []userData.User{}

	// CICLO FOR EACH FINCHE' ROW HA UN NEXT ASSEGNAMO AD U IL VALORE DEL SINGOLO OGGETTO SE C'È UN ERRORE AMEN ALTRIMENTI APPENS ALL'ARRAY
	for rows.Next() {
		var u userData.User
		if err := rows.Scan(&u.Email, &u.Username, &u.Password); err != nil {
			// Gestione dell'errore
			return 
		}
		users = append(users, u)
	}

	// IN RISPOSTA SETTIAMO HEADER CHE DEFINISCE CHE LA NOSTRA RISPOSTA E' UN JSON E CON ENCODER CONVERTIAMO ARRAY IN JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
