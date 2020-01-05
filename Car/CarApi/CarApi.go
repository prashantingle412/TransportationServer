package CarApi

import (
    "net/http"
    // "log"
    "encoding/json"
	"TransportationServer/CommonPackages/StructConfig"
	"TransportationServer/CommonPackages/Common"
	"TransportationServer/Car/CarDao"
	"github.com/gorilla/mux"
)
// Car Master
func AddCarMaster(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	args:= StructConfig.CarMaster{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil {
		Common.RespondWithError(w,http.StatusBadRequest,bodyErr.Error())
		return
	}
	err := CarDao.AddCarMaster(args)			
	if err != nil {
		Common.RespondWithError(w,http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Company Created successfully")
	}
}
func DisplayCarMaster(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	args,err := CarDao.GetCarMaster(params["id"])
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,args)
	}
}
func UpdateCarMaster(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	args:= StructConfig.CarMaster{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil {
		Common.RespondWithError(w,http.StatusBadRequest,bodyErr.Error())
	}
	err := CarDao.UpdateCarMaster(args)			
	if err != nil {
		Common.RespondWithError(w,http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Company Created successfully")
	}
}
func RemoveCarMaster(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := CarDao.DeleteCarMaster(params["id"])
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Car master Removed")
	}
}

// Car Details
func AddCompanyCar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	args:= StructConfig.CompanyCar{}
	bodyErr := json.NewDecoder(r.Body).Decode(&args)
	if bodyErr != nil {
		Common.RespondWithError(w,http.StatusBadRequest,bodyErr.Error())
	}
	err := CarDao.AddCompanyCar(args)			
	if err != nil {
		Common.RespondWithError(w,http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Company Car Created successfully")
	}
}
func DisplayCompanyCarInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	args,err := CarDao.GetCompanyCarInfo(params["id"])
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,args)
	}
}

func RemoveCarCompanyInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := CarDao.DeleteCarCompanyInfo(params["id"])
	if err != nil {
		Common.RespondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		Common.RespondWithJson(w,http.StatusOK,"Car master Removed")
	}
}