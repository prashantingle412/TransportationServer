
package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
	"TransportationServer/packages/RentalCompany"
)
func main() {
	r := mux.NewRouter()
	// company details api Endpoints
	r.HandleFunc("/companydetails",RentalCompany.CreateComapnyDetails).Methods("POST")
	r.HandleFunc("/companydetails/{id}", RentalCompany.DisplayComapnyDetails).Methods("GET")
	r.HandleFunc("/companydetails/{id}", RentalCompany.UpdateComapnyDetails).Methods("PUT")
	r.HandleFunc("/companydetails/{id}", RentalCompany.DeleteComapnyDetails).Methods("DELETE")

	// company location api 
	r.HandleFunc("/companylocation",RentalCompany.AddRentalCompanyLocation).Methods("POST")
	r.HandleFunc("/companylocation",RentalCompany.DisplayRentalCompanyLocation).Methods("GET")
    r.HandleFunc("/companylocation",RentalCompany.UpdateRentalCompanyLocation).Methods("PUT")
	r.HandleFunc("/companylocation",RentalCompany.RemoveLocation).Methods("DELETE")

	// for Auth 
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", RentalCompany.Login)
	s.Use(RentalCompany.IsuserExistMiddleware)
	
	log.Fatal(http.ListenAndServe(":4447", r))
	
}

/*
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
		log.Println(r.RequestURI)
		w.Write([]byte("this is midleware func"))
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}
func hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("with finalHandler")
	w.Write([]byte("hello from hello func"))
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("without Executing finalHandler  ")
	w.Write([]byte("without Executing finalHandler  "))
}
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is subrouter "))
}

func main() {  
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/product", ProductsHandler)
	r.HandleFunc("/final", final)
	s.HandleFunc("/", hello)
	s.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8000", r))
}
*/