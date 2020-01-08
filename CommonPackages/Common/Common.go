package Common

import (
	"log"
	"net/http"
	// "log"
	"crypto/md5"
	"encoding/json"
	"fmt"

	"context"

	"github.com/go-playground/validator"

	"github.com/google/uuid"
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

type ContextKey string

const ContextKeyRequestID ContextKey = "requestID"

// func for assign rquest id
func AssignRequestID(ctx context.Context) context.Context {

	reqID := uuid.New()

	return context.WithValue(ctx, ContextKeyRequestID, reqID.String())
}

// get request id
func GetRequestID(ctx context.Context) string {

	reqID := ctx.Value(ContextKeyRequestID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}
func ReqIDMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		r = r.WithContext(AssignRequestID(ctx))

		log.Println("Incomming request %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, GetRequestID(ctx))

		next.ServeHTTP(w, r)

		log.Println("Finished handling http req. %s", GetRequestID(ctx))
	})
}
