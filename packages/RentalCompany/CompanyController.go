package RentalCompany
import (
    "net/http"
    // "log"
    "encoding/json"
	"fmt"
    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"TransportationServer/packages/Common"
	// "MuxApi/packages/TokenManager"
)
var sess *mgo.Session
var collection *mgo.Collection
var sessUValName string

// Rest api for Create,update,Show, and Delete Company details
func CreateComapnyDetails(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()	
	args := Common.Company{}
    _ = json.NewDecoder(r.Body).Decode(&args)
	collection = setCollection("j1_db","company_collection")
	str := &Common.Company{Id:bson.ObjectId(bson.NewObjectId()).Hex(),CreatedOn:time.Now().UnixNano() / (int64(time.Millisecond)),CompanyName:args.CompanyName,CompanyRegistrationNumber:args.CompanyRegistrationNumber,CompanyEmail:args.CompanyEmail,PhoneNumber:args.PhoneNumber,MobileNumber:args.MobileNumber,UserId:args.UserId,Password:args.Password}
	err := collection.Insert(str)
	if err != nil {
		fmt.Println("error in inserting",err)
	}else {
		userArgs := Common.UserInstance{}
		collection = setCollection("j1_db","userInstance_collection")
		UserIntanceStr := &Common.UserInstance{Id:bson.ObjectId(bson.NewObjectId()).Hex(),UserAddedOn:time.Now().UnixNano() / (int64(time.Millisecond)),UserEmail:userArgs.UserEmail,UserName:userArgs.UserName}
		err2 := collection.Insert(UserIntanceStr)			
		if err2 != nil {
			respondWithError(w,http.StatusBadRequest,err2.Error())
		}else{
			respondWithJson(w,http.StatusOK,args)
		}
	}
}

func DisplayComapnyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    fmt.Println("vars are ",params["id"] )
	collection = setCollection("j1_db","company_collection")
    args := Common.Company{}
    err := collection.Find(bson.M{"_id":params["id"]}).One(&args)
	if err != nil {
		respondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		respondWithJson(w,http.StatusOK,args)
	}
}

func UpdateComapnyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("vars are ",params["id"] )
	args := Common.Company{}
    _ = json.NewDecoder(r.Body).Decode(&args)
    collection = setCollection("j1_db","company_collection")
    err := collection.Update(bson.M{"_id":params["id"]},bson.M{"$set":bson.M{"company_name":args.CompanyName,"company_registration_number":args.CompanyRegistrationNumber,"company_email":args.CompanyEmail,"phone_number":args.PhoneNumber,"mobile_number":args.MobileNumber,"user_id":args.UserId,"password":args.Password}})
	if err != nil {
		respondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		respondWithJson(w,http.StatusOK,args)
	}
}
func DeleteComapnyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    collection = setCollection("j1_db","company_collection")
    err := collection.Remove(bson.M{"_id":params["id"]})
	if err != nil {
		respondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		respondWithJson(w,http.StatusOK,"Company Details Successfully")
	}
}

// for add,show,update,delete location detail of rental company
func AddRentalCompanyLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	args:= make(map[string]interface{})
	_ = json.NewDecoder(r.Body).Decode(&args)
	collection = setCollection("j1_db","company_location")
	insertQuery := &Common.CompanyLocation{Id:bson.ObjectId(bson.NewObjectId()).Hex(),MobileNumber:args["mobileNumber"].(string),LocationAddedOn:time.Now().UnixNano() / (int64(time.Millisecond)),LocationName:args["locationName"].(string),CoodinatesX:args["coordinatesX"].(string),CoordinatesY:args["coordinatesY"].(string)}
	err := collection.Insert(insertQuery)
	if err != nil {
		respondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		respondWithJson(w,http.StatusOK,args)
	}
}

func DisplayRentalCompanyLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	params := mux.Vars(r)
	collection = setCollection("j1_db","company_location")
	companyLocationStr := Common.CompanyLocation{}
	err := collection.Find(bson.M{"_id":params["id"]}).One(&companyLocationStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		respondWithJson(w,http.StatusOK,companyLocationStr)
	}
}

func UpdateRentalCompanyLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)	
	args := make(map[string]interface{})
	_ = json.NewDecoder(r.Body).Decode(&args)
	collection = setCollection("j1_db","company_location")
	err := collection.Update(bson.M{"_id":params["id"]},bson.M{"$set":bson.M{"mobile_number":args["mobileNumber"].(string),"location_name":args["locationName"].(string),"coordinates_x":args["coordinatesX"].(string),"coordinates_y":args["coordinatesY"].(string)}})
	if err != nil {
		respondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		respondWithJson(w,http.StatusOK,args)
	}
}

func RemoveLocation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	params := mux.Vars(r)	
	collection = setCollection("j1_db","company_location")
	err := collection.Remove(bson.M{"_id":params["Id"]})
	if err != nil {
		respondWithError(w, http.StatusBadRequest,err.Error())
	}else{
		respondWithJson(w,http.StatusOK,"Company Location Removed Successfully")
	}
}
// Api is in under working 
func Login(w http.ResponseWriter, r *http.Request){
	respondWithJson(w,http.StatusOK,"welcome login successfull .....user authenticated ")
	// collection = setCollection("")
}

func IsuserExistMiddleware( next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	
	// m := Common.UserInstance{}
	m := make(map[string]interface{})
	_ = json.NewDecoder(r.Body).Decode(&m)	
	collection = setCollection("j1_db", "company_collection")
	// pwd := StringMd5(password)
	usr, err := collection.Find(bson.M{"company_email": m["email"].(string), "password": m["password"].
	(string)}).Count()
	if err != nil {
		if err.Error() == "not found" {
			respondWithError(w, http.StatusBadRequest,"user does not exists")
			// return false, nil
		} else {
			respondWithError(w, http.StatusBadRequest,"user does not exists")
			// return false, err
		}
	} else if usr < 1 {
		respondWithError(w, http.StatusBadRequest,"user does not exists")
	} else {
		next.ServeHTTP(w, r)
	  }
	})
}

//Helper Functions  

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, result interface{}) {
    response, _ := json.Marshal(result)
    fmt.Println("response sis ",response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func setCollection(dbName string, collectionName string) *mgo.Collection {
	if sess == nil {
		fmt.Println("Not connected... Connecting to Mongo")
		sess = GetConnected()
	}
	collection = sess.DB(dbName).C(collectionName)
	return collection
}

func GetConnected() *mgo.Session {
	dialInfo, err := mgo.ParseURL("mongodb://localhost:27017")
	dialInfo.Direct = true
	dialInfo.FailFast = true
	dialInfo.Database = "j1_db"
	dialInfo.Username = "root"
	dialInfo.Password = "tiger"
	sess, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("Can't connect to mongo, go error %v\n", err)
		panic(err)
	} else {
		return sess
		defer sess.Close()
	}
	return sess
}
