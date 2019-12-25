package Customer
import (
	"fmt"
	"net/http"
)
func Hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("customer ")
}