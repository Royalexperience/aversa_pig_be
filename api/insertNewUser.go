package apis

import "net/http"
import "encoding/json"
import "database/sql"
import "aversan/BE/types"
import "aversan/BE/utils"

func InsertNewData (w http.ResponseWriter, r *http.Request, db *sql.DB) {
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user userData.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	// Controlla se il corpo della richiesta è valido
	// VA FATTO MA FE MA PER ORA Lo LASCIO
	if (checkData(user, w)){
		return
	}

	sqlStatement := utilsConstants.QUERY_INSERT_NEW_DATA
	_, err = db.Exec(sqlStatement, user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "[Error 500] Error trying to execute query ", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "User created successfully"}
	responseBytes, err := json.Marshal(response)
	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBytes)

}

func checkData ( user userData.User, w http.ResponseWriter) bool {
		var errorMessage string
		var errorHappened bool = false
		if(user.Email == ""){
			errorHappened = true 
			errorMessage = "EMAIL must not be null or empty"
		}
		if(user.Username == ""){
			errorHappened = true 
			errorMessage = "USERNAME must not be null or empty"
		}
		
		if(user.Password == ""){
			errorHappened = true 
			errorMessage = "PASSWORD must not be null or empty"
		}
		response := map[string]string{
            "error": "[ERROR 400] Bad Request",
            "message": errorMessage,
        }
		responseBytes, err := json.Marshal(response)
        if err != nil {
            http.Error(w, "Error Parsing Response", http.StatusInternalServerError)
            return errorHappened
        }
		if (errorHappened){
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(responseBytes)
		}
		
		return errorHappened
}