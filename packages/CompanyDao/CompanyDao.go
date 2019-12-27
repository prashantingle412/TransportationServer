package CompanyDao
import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"TransportationServer/packages/StructConfig"
	"TransportationServer/packages/DbConfig"
	// "log"
)

func AddCompanyDetails(args StructConfig.Company) error{
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_collection")
	str := &StructConfig.Company{Id:bson.ObjectId(bson.NewObjectId()).Hex(),CreatedOn:time.Now().UnixNano() / (int64(time.Millisecond)),CompanyName:args.CompanyName,CompanyRegistrationNumber:args.CompanyRegistrationNumber,Email:args.Email,PhoneNumber:args.PhoneNumber,MobileNumber:args.MobileNumber,UserId:args.UserId,Password:args.Password,UserRole:args.UserRole}
	err := DbConfig.Collection.Insert(str)
	if err != nil {
		return err
	}else {
		DbConfig.Collection = DbConfig.SetCollection("transportation_db","userInstance_collection")
		UserIntanceStr := &StructConfig.UserInstance{MobileNumber:args.MobileNumber,Id:bson.ObjectId(bson.NewObjectId()).Hex(),UserAddedOn:time.Now().UnixNano() / (int64(time.Millisecond)),UserEmail:args.Email,UserName:args.CompanyName,UserRole:args.UserRole}
		err2 := DbConfig.Collection.Insert(UserIntanceStr)		
		if err2 != nil {
			return err
		}	
	}
	return nil
}
func GetCompanyDetails(Id string ) (StructConfig.Company,error) {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_collection")
    args := StructConfig.Company{}
	err := DbConfig.Collection.Find(bson.M{"_id":Id}).One(&args)
	return args,err
}
func PutCompanyDeails(args StructConfig.Company) error{
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_collection")
	err := DbConfig.Collection.Update(bson.M{"_id":args.Id},bson.M{"$set":bson.M{"company_name":args.CompanyName,"company_registration_number":args.CompanyRegistrationNumber,"email":args.Email,"phone_number":args.PhoneNumber,"mobile_number":args.MobileNumber,"user_id":args.UserId,"password":args.Password,"user_role":args.UserRole}})
	return err
}

func RemoveCompanyInfo(Id string) error{
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_collection")
	err := DbConfig.Collection.Remove(bson.M{"_id":Id})
	return err
}

func AddCLocation(args StructConfig.CompanyLocation) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_location_collection")
	str := &StructConfig.CompanyLocation{Id:bson.ObjectId(bson.NewObjectId()).Hex(),MobileNumber:args.MobileNumber,LocationAddedOn:time.Now().UnixNano() / (int64(time.Millisecond)),LocationName:args.LocationName,CoodinatesX:args.CoodinatesX,CoordinatesY:args.CoordinatesY}
	err := DbConfig.Collection.Insert(str)
	return err
}
func ShowCLocation(Id string) (StructConfig.CompanyLocation, error) {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_location_collection")
	companyLocationStr := StructConfig.CompanyLocation{}
	err := DbConfig.Collection.Find(bson.M{"_id":Id}).One(&companyLocationStr)
	return companyLocationStr,err
}

func UpdateCLocation(args StructConfig.CompanyLocation) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_location_collection")
	err := DbConfig.Collection.Update(bson.M{"_id":args.Id},bson.M{"$set":bson.M{"mobile_number":args.MobileNumber,"location_name":args.LocationName,"coordinates_x":args.CoodinatesX,"coordinates_y":args.CoordinatesY}})
	return err
}

func RemoveCLocation(Id string) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db","company_location_collection")
	err := DbConfig.Collection.Remove(bson.M{"_id":Id})
	return err
}
