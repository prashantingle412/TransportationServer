
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
	r.HandleFunc("/companylocation/{id}",RentalCompany.DisplayRentalCompanyLocation).Methods("GET")
    r.HandleFunc("/companylocation/{id}",RentalCompany.UpdateRentalCompanyLocation).Methods("PUT")
	r.HandleFunc("/companylocation/{id}",RentalCompany.RemoveLocation).Methods("DELETE")

	// for Auth 
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", RentalCompany.Login)
	s.Use(RentalCompany.IsuserExistMiddleware)
	
	log.Fatal(http.ListenAndServe(":4447", r))
	
}