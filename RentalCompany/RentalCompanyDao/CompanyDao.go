package CompanyDao

import (
	"TransportationServer/CommonPackages/DbConfig"
	"TransportationServer/CommonPackages/StructConfig"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	// "encoding/json"
	"TransportationServer/CommonPackages/Common"
)

func AddCompanyDetails(args StructConfig.Company) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_collection")
	Password := Common.StringMd5(args.Password)
	str := &StructConfig.Company{Id: bson.ObjectId(bson.NewObjectId()).Hex(), CreatedOn: time.Now().UnixNano() / (int64(time.Millisecond)), CompanyName: args.CompanyName, CompanyRegistrationNumber: args.CompanyRegistrationNumber, Email: args.Email, PhoneNumber: args.PhoneNumber, MobileNumber: args.MobileNumber, UserId: args.UserId, Password: Password, UserRole: args.UserRole}
	valdateErr := Common.ValidateStructFeild(str)
	if valdateErr != nil {
		fmt.Println("error msg:", valdateErr.Error())
		return valdateErr
	}
	err := DbConfig.Collection.Insert(str)
	if err != nil {
		return err
	} else {
		DbConfig.Collection = DbConfig.SetCollection("transportation_db", "userInstance_collection")
		UserIntanceStr := &StructConfig.UserInstance{MobileNumber: args.MobileNumber, Id: bson.ObjectId(bson.NewObjectId()).Hex(), UserAddedOn: time.Now().UnixNano() / (int64(time.Millisecond)), UserEmail: args.Email, UserName: args.CompanyName, UserRole: args.UserRole, Password: Password}
		ValidateUserErr := Common.ValidateStructFeild(UserIntanceStr)
		if ValidateUserErr != nil {
			return ValidateUserErr
		}
		err2 := DbConfig.Collection.Insert(UserIntanceStr)
		if err2 != nil {
			return err
		}
	}
	return err
}
func GetCompanyDetails(Id string) (StructConfig.Company, error) {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_collection")
	args := StructConfig.Company{}
	err := DbConfig.Collection.Find(bson.M{"_id": Id}).One(&args)
	return args, err
}
func PutCompanyDeails(args StructConfig.Company) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_collection")
	err := DbConfig.Collection.Update(bson.M{"_id": args.Id}, bson.M{"$set": bson.M{"company_name": args.CompanyName, "company_registration_number": args.CompanyRegistrationNumber, "email": args.Email, "phone_number": args.PhoneNumber, "mobile_number": args.MobileNumber, "user_id": args.UserId, "password": Common.StringMd5(args.Password), "user_role": args.UserRole}})
	return err
}

func RemoveCompanyInfo(Id string) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_collection")
	err := DbConfig.Collection.Remove(bson.M{"_id": Id})
	return err
}

func AddCLocation(args StructConfig.CompanyLocation) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_location_collection")
	str := &StructConfig.CompanyLocation{Id: bson.ObjectId(bson.NewObjectId()).Hex(), MobileNumber: args.MobileNumber, LocationAddedOn: time.Now().UnixNano() / (int64(time.Millisecond)), LocationName: args.LocationName, CoodinatesX: args.CoodinatesX, CoordinatesY: args.CoordinatesY}
	validateErr := Common.ValidateStructFeild(str)
	if validateErr != nil {
		return validateErr
	}
	err := DbConfig.Collection.Insert(str)
	return err
}
func ShowCLocation(Id string) (StructConfig.CompanyLocation, error) {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_location_collection")
	companyLocationStr := StructConfig.CompanyLocation{}
	err := DbConfig.Collection.Find(bson.M{"_id": Id}).One(&companyLocationStr)
	return companyLocationStr, err
}

func UpdateCLocation(args StructConfig.CompanyLocation) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_location_collection")
	err := DbConfig.Collection.Update(bson.M{"_id": args.Id}, bson.M{"$set": bson.M{"mobile_number": args.MobileNumber, "location_name": args.LocationName, "coordinates_x": args.CoodinatesX, "coordinates_y": args.CoordinatesY}})
	return err
}

func RemoveCLocation(Id string) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_location_collection")
	err := DbConfig.Collection.Remove(bson.M{"_id": Id})
	return err
}
func IsUserExist(args StructConfig.UserInstance) (bool, error) {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_collection")
	// pwd := StringMd5(password)
	usr, err := DbConfig.Collection.Find(bson.M{"company_email": args.UserEmail, "password": Common.StringMd5(args.Password)}).Count()
	if err != nil {
		if err.Error() == "not found" {
			return false, err
		} else {
			return false, err
		}
	} else if usr == 1 {
		return true, nil
	} else {
		return false, err
	}
}
func CheckRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("user_role")
		email := r.Header.Get("user_email")
		DbConfig.Collection = DbConfig.SetCollection("transportation_db", "userInstance_collection")
		admin, err := DbConfig.Collection.Find(bson.M{"user_email": email, "user_role": role}).Count()
		if err != nil {
			if err.Error() == "not found" {
				Common.RespondWithError(w, http.StatusBadRequest, "You are not Authenticate to perform Action")
			} else {
				Common.RespondWithError(w, http.StatusBadRequest, "You are not Authenticate to perform Action")
			}
		} else if admin == 1 {
			log.Println("count is ", admin)
			next.ServeHTTP(w, r)
		} else {
			Common.RespondWithError(w, http.StatusBadRequest, "You are Authenticate to perform Action")
		}
	})
}
