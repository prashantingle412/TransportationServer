package CarDao

import (
	"TransportationServer/CommonPackages/Common"
	"TransportationServer/CommonPackages/DbConfig"
	"TransportationServer/CommonPackages/StructConfig"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Car started
func AddCarMaster(args StructConfig.CarMaster) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "carmaster_collection")
	str := &StructConfig.CarMaster{Id: bson.ObjectId(bson.NewObjectId()).Hex(), AddedOn: time.Now().UnixNano() / (int64(time.Millisecond)), CarCategory: args.CarCategory, CarMake: args.CarMake}
	validateErr := Common.ValidateStructFeild(str)
	if validateErr != nil {
		return validateErr
	}
	err := DbConfig.Collection.Insert(str)
	return err
}
func GetCarMaster(Id string) (StructConfig.CarMaster, error) {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "carmaster_collection")
	args := StructConfig.CarMaster{}
	err := DbConfig.Collection.Find(bson.M{"_id": Id}).One(&args)
	return args, err
}
func UpdateCarMaster(args StructConfig.CarMaster) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "carmaster_collection")
	err := DbConfig.Collection.Update(bson.M{"_id": args.Id}, bson.M{"$set": bson.M{"car_category": args.CarCategory, "car_make": args.CarMake}})
	return err
}

func DeleteCarMaster(Id string) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "carster_collection")
	err := DbConfig.Collection.Remove(bson.M{"_id": Id})
	return err
}
func AddCompanyCar(args StructConfig.CompanyCar) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_carinfo_collection")
	CarFetureArgs := args.CarFetures
	CarInsuranceInfoArgs := args.CarInsuranceInfo
	CarInsuranceInfoStr := StructConfig.CarInsuranceInfo{InsuranceCompanyName: CarInsuranceInfoArgs.InsuranceCompanyName, InsuranceProvider: CarInsuranceInfoArgs.InsuranceProvider, PolicyNumber: CarInsuranceInfoArgs.PolicyNumber, Website: CarInsuranceInfoArgs.Website}
	CarFeatureStr := StructConfig.CarFetures{WindowType: CarFetureArgs.WindowType, SoundType: CarFetureArgs.SoundType, LockingType: CarFetureArgs.LockingType, ParkingCensor: CarFetureArgs.ParkingCensor, AirBags: CarFetureArgs.AirBags, GearType: CarFetureArgs.GearType, Camera: CarFetureArgs.Camera, DoorLock: CarFetureArgs.DoorLock, CarAc: CarFetureArgs.CarAc, PowerStreeing: CarFetureArgs.PowerStreeing}
	str := &StructConfig.CompanyCar{Id: bson.ObjectId(bson.NewObjectId()).Hex(), CompanyCarAddedOn: time.Now().UnixNano() / (int64(time.Millisecond)), CarCategoryId: args.CarCategoryId, CarNumberPLate: args.CarNumberPLate, CarColor: args.CarColor, CarModel: args.CarModel, CarYear: args.CarYear, CarMilage: args.CarMilage, CarAddress: args.CarAddress, CarName: args.CarName, CarValue: args.CarValue, CarInsuranceInfo: CarInsuranceInfoStr, CarFetures: CarFeatureStr}
	validateErr := Common.ValidateStructFeild(str)
	if validateErr != nil {
		return validateErr
	}
	err := DbConfig.Collection.Insert(str)
	return err
}

func GetCompanyCarInfo(Id string) (StructConfig.CompanyCar, error) {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_carinfo_collection")
	args := StructConfig.CompanyCar{}
	err := DbConfig.Collection.Find(bson.M{"_id": Id}).One(&args)
	return args, err
}
func DeleteCarCompanyInfo(Id string) error {
	DbConfig.Collection = DbConfig.SetCollection("transportation_db", "company_carinfo_collection")
	err := DbConfig.Collection.Remove(bson.M{"_id": Id})
	return err
}
