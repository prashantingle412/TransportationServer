
package main

import (
    "net/http"
	"log"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"TransportationServer/packages/RentalCompanyApi"
)
func main() {
	r := mux.NewRouter()
	// company details api Endpoints
	r.HandleFunc("/companydetails",RentalCompanyApi.CreateComapnyDetails).Methods("POST")
	r.HandleFunc("/companydetails/{id}", RentalCompanyApi.DisplayComapnyDetails).Methods("GET")
	r.HandleFunc("/companydetails/{id}", RentalCompanyApi.UpdateComapnyDetails).Methods("PUT")
	r.HandleFunc("/companydetails/{id}", RentalCompanyApi.DeleteComapnyDetails).Methods("DELETE")

	// company location api 
	r.HandleFunc("/companylocation",RentalCompanyApi.AddRentalCompanyLocation).Methods("POST")
	r.HandleFunc("/companylocation/{id}",RentalCompanyApi.DisplayRentalCompanyLocation).Methods("GET")
    r.HandleFunc("/companylocation/{id}",RentalCompanyApi.UpdateRentalCompanyLocation).Methods("PUT")
	r.HandleFunc("/companylocation/{id}",RentalCompanyApi.RemoveLocation).Methods("DELETE")

	// for Auth 
	// s := r.PathPrefix("/auth").Subrouter()
	// s.HandleFunc("/login", RentalCompanyApi.Login)
	// s.Use(RentalCompanyApi.IsuserExistMiddleware)
	
	log.Fatal(http.ListenAndServe(":4447",handlers.LoggingHandler(os.Stdout, r)))
}