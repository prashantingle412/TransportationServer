package Common

import (
	"net/http"
	// "log"
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator"
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
func ValidateStructFeild(str interface{}) error {
	valdateErr := validator.New().Struct(str)
	if valdateErr != nil {
		fmt.Println("validation func", valdateErr)
		return valdateErr
	}
	return nil
}
