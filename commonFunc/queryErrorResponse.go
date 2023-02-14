package utilityFunc
import "net/http"
import "aversan/BE/utils"
import "encoding/json"
func QueryErrorResponse (errorResponse error,w http.ResponseWriter, queryErrorMessage string) {
		response500 := map[string]string{
			"error":   utilsConstants.ERROR_500,
			"message": queryErrorMessage,
		}
		responseBytes, err := json.Marshal(response500)
		ErrorParsingResponse(err,w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responseBytes)
}