package utilityFunc

import "net/http"
import "encoding/json"
import "aversan/BE/utils"

func CheckMethod(r *http.Request, w http.ResponseWriter,method string) bool {
	
	if r.Method != method {
		response405 := map[string]string{
			"error":   utilsConstants.ERROR_405,
			"message": r.Method + " " + utilsConstants.METHOD_NOT_ALLOWED,
		}
		responseBytes, err := json.Marshal(response405)
		ErrorParsingResponse(err,w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseBytes)
		return true
	}
	return false
}

