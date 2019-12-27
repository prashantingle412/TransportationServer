package Common
import (
	"net/http"
	"log"
	"encoding/json"
)
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, result interface{}) {
    response, _ := json.Marshal(result)
    log.Println("response sis ",response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
