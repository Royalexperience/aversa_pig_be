package utilityFunc
import "net/http"
import "aversan/BE/utils"

func ErrorParsingResponse (errorResponse error,w http.ResponseWriter) {
	if errorResponse != nil {
		http.Error(w, utilsConstants.ERROR_PARSING_RESPONSE, http.StatusInternalServerError)
	}
}

