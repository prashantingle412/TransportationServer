package RentalCompanyApi
import (
    "net/http"
    // "log"
    "encoding/json"
    "github.com/gorilla/mux"
    // "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "time"
	"TransportationServer/packages/StructConfig"
	"TransportationServer/packages/CompanyDao"
	"TransportationServer/packages/Common"
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
	_ = json.NewDecoder(r.Body).Decode(&args)
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
/*
// Api is in under working 
func Login(w http.ResponseWriter, r *http.Request){
	Common.RespondWithJson(w,http.StatusOK,"welcome login successfull .....user authenticated ")
	// collection = setCollection("")
}

func IsuserExistMiddleware( next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	// m := StructConfig.UserInstance{}
	m := make(map[string]interface{})
	_ = json.NewDecoder(r.Body).Decode(&m)	
	collection = setCollection("transportation_db", "company_collection")
	// pwd := StringMd5(password)
	usr, err := collection.Find(bson.M{"company_email": m["email"].(string), "password": m["password"].
	(string)}).Count()
	if err != nil {
		if err.Error() == "not found" {
			Common.RespondWithError(w, http.StatusBadRequest,"user does not exists")
			// return false, nil
		} else {
			Common.RespondWithError(w, http.StatusBadRequest,"user does not exists")
			// return false, err
		}
	} else if usr < 1 {
		Common.RespondWithError(w, http.StatusBadRequest,"user does not exists")
	} else {
		next.ServeHTTP(w, r)
	  }
	})
}
*/

//Helper Functions  

