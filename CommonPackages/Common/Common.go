package Common
import (
	"net/http"
	// "log"
	"encoding/json"
	"crypto/md5"
	"fmt"
)
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, result interface{}) {
    response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func StringMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}