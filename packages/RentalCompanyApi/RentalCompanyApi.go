package RentalCompanyApi
import (
    "net/http"
    // "log"
    "encoding/json"
    "github.com/gorilla/mux"
	"TransportationServer/packages/StructConfig"
	"TransportationServer/packages/CompanyDao"
	"TransportationServer/packages/Common"
	"TransportationServer/packages/TokenManager"
	jwt "github.com/dgrijalva/jwt-go"
	
	"fmt"
)

// Rest api for Create,update,Show, and Delete Company details
func CreateComapnyDetails(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()	
	args := StructConfig.Company{}
    bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil {
		Common.RespondWithError(w,http.StatusBadRequest,bodyErr.Error())		
		return
	}
	err := CompanyDao.AddCompanyDetails(args)			
	if err != nil {
		Common.RespondWithError(w,http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Company Created successfully")
	}
}

func DisplayComapnyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	args,err := CompanyDao.GetCompanyDetails(params["id"])
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,args)
	}
}

func UpdateComapnyDetails(w http.ResponseWriter, r *http.Request) {
	args := StructConfig.Company{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil {
		Common.RespondWithError(w, http.StatusBadRequest,bodyErr.Error())
		return
	}
	err := CompanyDao.PutCompanyDeails(args)
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"company updated successfully")
	}
}
func DeleteComapnyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    err := CompanyDao.RemoveCompanyInfo(params["id"])
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Company Details Successfully")
	}
}

// for add,show,update,delete location detail of rental company
func AddRentalCompanyLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	args:= StructConfig.CompanyLocation{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil {
		Common.RespondWithError(w, http.StatusBadRequest,bodyErr.Error())
	}
	err := CompanyDao.AddCLocation(args)
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,args)
	}
}

func DisplayRentalCompanyLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	params := mux.Vars(r)
	companyLocationStr,err := CompanyDao.ShowCLocation(params["id"]) 	
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,companyLocationStr)
	}
}

func UpdateRentalCompanyLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	args := StructConfig.CompanyLocation{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil{
		Common.RespondWithError(w, http.StatusBadRequest,bodyErr.Error())
	}
	err :=CompanyDao.UpdateCLocation(args)
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,args)
	}
}

func RemoveLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	params := mux.Vars(r)	
	err :=CompanyDao.RemoveCLocation(params["id"])
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Company Location Removed Successfully")
	}
}

// Api is in under working 
func Login(w http.ResponseWriter, r *http.Request){	
	args := StructConfig.UserInstance{}	
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil {
		Common.RespondWithError(w,http.StatusBadRequest ,bodyErr.Error())		
	}
	if _,userexistErr := CompanyDao.IsUserExist(args); userexistErr != nil{
		Common.RespondWithError(w,http.StatusBadRequest ,userexistErr.Error())
		return
	}
	if token,err := TokenManager.GenerateToken(args); err != nil{
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else {
		Common.RespondWithJson(w,http.StatusOK,token)
	}
}
func IsAuthorized(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return []byte("write_some_secret_key_here"), nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                next.ServeHTTP(w, r)
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
	})
}
func Welcome(w http.ResponseWriter, r *http.Request) {
	Common.RespondWithJson(w,http.StatusOK,"welcome to secret part")
}
/*
// register user/ customer/ admin
func NewRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	args := StructConfig.Employee{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil{
		Common.RespondWithError(w, http.StatusBadRequest,bodyErr.Error())
	}
	err :=CompanyDao.AddEmployee(args)
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())		
	}else {
		Common.RespondWithJson(w, http.StatusOK,"employee created successfully")				
	}
}
*/

//  Car APIs started
// func AddCarDetails(w http.ResponseWriter, r *http.Request) {
// 	args := StructConfig.{}	
// 	bodyErr := json.NewDecoder(r.Body).Decode(&args)
// }