/*
package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

// "fmt"
// "encoding/json"

type User struct {
	Name string `json:"name" validate:"required"`
	City string `json:"city" validate:"required"`
}

func main() {
	userObj := User{
		Name: "",
		City: "pune",
	}
	validate := validator.New()
	err := validate.Struct(userObj)
	if err != nil {
		fmt.Println("error in validate user",err.Error())
	}else {
		fmt.Println("validated")
	}
}

*/
package main

import (
	"TransportationServer/Car/CarApi"
	"TransportationServer/RentalCompany/RentalCompanyApi"
	CompanyDao "TransportationServer/RentalCompany/RentalCompanyDao"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", RentalCompanyApi.Login).Methods("GET")
	// sub router for generating jwt token for user

	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/welcome", RentalCompanyApi.Welcome)
	s.Use(RentalCompanyApi.IsAuthorized)
	// api For create company with Defoult Admin
	r.HandleFunc("/createcompany", RentalCompanyApi.CreateComapnyDetails).Methods("POST")
	// We check if roleID is Admin or not and then do other operations
	a := r.PathPrefix("/admin").Subrouter()
	a.HandleFunc("/companydetails/{id}", RentalCompanyApi.DisplayComapnyDetails).Methods("GET")
	a.HandleFunc("/companydetails/{id}", RentalCompanyApi.UpdateComapnyDetails).Methods("PUT")
	a.HandleFunc("/companydetails/{id}", RentalCompanyApi.DeleteComapnyDetails).Methods("DELETE")
	// company location api
	a.HandleFunc("/companylocation", RentalCompanyApi.AddRentalCompanyLocation).Methods("POST")
	a.HandleFunc("/companylocation/{id}", RentalCompanyApi.DisplayRentalCompanyLocation).Methods("GET")
	a.HandleFunc("/companylocation/{id}", RentalCompanyApi.UpdateRentalCompanyLocation).Methods("PUT")
	a.HandleFunc("/companylocation/{id}", RentalCompanyApi.RemoveLocation).Methods("DELETE")

	a.HandleFunc("/carmaster", CarApi.AddCarMaster).Methods("POST")
	a.HandleFunc("/carmaster/{id}", CarApi.DisplayCarMaster).Methods("GET")
	a.HandleFunc("/carmaster/{id}", CarApi.UpdateCarMaster).Methods("PUT")
	a.HandleFunc("/carmaster/{id}", CarApi.RemoveCarMaster).Methods("DELETE")
	//For car informations
	a.HandleFunc("/addcar", CarApi.AddCompanyCar).Methods("DELETE")
	a.HandleFunc("/carmaster/{id}", CarApi.DisplayCompanyCarInfo).Methods("GET")
	a.HandleFunc("/carmaster/{id}", CarApi.RemoveCarCompanyInfo).Methods("DELETE")

	a.Use(CompanyDao.CheckRole)

	log.Fatal(http.ListenAndServe(":4447", handlers.LoggingHandler(os.Stdout, r)))
}
