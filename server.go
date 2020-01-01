
package main

import (
    "net/http"
	"log"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"TransportationServer/packages/RentalCompanyApi"
	"TransportationServer/packages/CompanyDao"
)
func main() {
	r := mux.NewRouter()
	// api for register user customer and admin
	// r.HandleFunc("/register",RentalCompanyApi.NewRegister).Methods("POST")
	// company details api sub router Endpoints
	a := r.PathPrefix("/admin").Subrouter()
	a.HandleFunc("/companydetails",RentalCompanyApi.CreateComapnyDetails).Methods("POST")
	a.HandleFunc("/companydetails/{id}", RentalCompanyApi.DisplayComapnyDetails).Methods("GET")
	a.HandleFunc("/companydetails/{id}", RentalCompanyApi.UpdateComapnyDetails).Methods("PUT")
	a.HandleFunc("/companydetails/{id}", RentalCompanyApi.DeleteComapnyDetails).Methods("DELETE")
	// company location api 
	a.HandleFunc("/companylocation",RentalCompanyApi.AddRentalCompanyLocation).Methods("POST")
	a.HandleFunc("/companylocation/{id}",RentalCompanyApi.DisplayRentalCompanyLocation).Methods("GET")
    a.HandleFunc("/companylocation/{id}",RentalCompanyApi.UpdateRentalCompanyLocation).Methods("PUT")
	a.HandleFunc("/companylocation/{id}",RentalCompanyApi.RemoveLocation).Methods("DELETE")
	a.Use(CompanyDao.CheckRole)

	// for Auth 
	r.HandleFunc("/login",RentalCompanyApi.Login).Methods("GET")
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/welcome", RentalCompanyApi.Welcome)
	s.Use(RentalCompanyApi.IsAuthorized)
	
	log.Fatal(http.ListenAndServe(":4447",handlers.LoggingHandler(os.Stdout, r)))

}

