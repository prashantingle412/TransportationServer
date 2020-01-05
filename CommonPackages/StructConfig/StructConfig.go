package StructConfig
import jwt "github.com/dgrijalva/jwt-go"
type Company struct {
    Id string `json:"_id" bson:"_id"`
	CreatedOn int64 `json:"created_on" bson:"created_on"`
	CompanyName string `json:"company_name" bson:"company_name"`
    CompanyRegistrationNumber string `json:"company_registration_number" bson:"company_registration_number"`
    Email string `json:"email" bson:"email"`
    PhoneNumber string `json:"phone_number" bson:"phone_number"`
    MobileNumber string `json:"mobile_number" bson:"mobile_number"`
    UserId string       `json:"user_id" bson:"user_id"`
    Password string `json:"password" bson:"password"`
	UserRole string `json:"user_role" bson:"user_role"`
}
type CompanyLocation struct {
    Id string `json:"_id" bson:"_id"`
    MobileNumber string `json:"mobile_number" bson:"mobile_number"`
    LocationAddedOn int64`json:"location_added_on" bson:"location_added_on"`
    LocationName string `json:"location_name" bson:"location_name"`
    CoodinatesX string `json:"coordinates_x" bson:"coordinates_x"`
    CoordinatesY string `json:"coordinates_y" bson:"coordinates_y"`
}
type UserInstance struct {
    MobileNumber string `json:"mobile_number" bson:"mobile_number"` // common in comapny and user
    Id string `json:"_id" bson:"_id"`
    UserAddedOn int64 `json:"user_added_on" bson:"user_added_on"`
    UserEmail string `json:"user_email" bson:"user_email"`
    UserName string `json:"user_name" bson:"user_name"`    
    UserRole string `json:"user_role" bson:"user_role"`
    Password string `json:"password" bson:"password"`
}
type JwtStruct struct {
    UserEmail string `json:"user_email"`
    UserName string `json:"user_name"`   
    jwt.StandardClaims 
}
type CarMaster struct {
    Id string `json:"_id" bson:"_id"`
    AddedOn int64 `json:"added_on" bson:"added_on"`
    CarCategory string `json:"car_category" bson:"car_category"`
    CarMake string`json:"car_make" bson:"car_make"`  
}
type CompanyCar struct {
    Id string `json:"_id" bson:"_id"`
    CompanyCarAddedOn int64 `json:"company_car_addedon" bson:"company_car_addedon"` 
    CarCategoryId string `json:"car_category_id" bson:"car_category_id"`
    CarNumberPLate string `json:"car_number_plate" bson:"car_number_plate"`
    CarColor string `json:"car_color" bson:"car_color"`
    CarModel string `json:"car_model" bson:"car_model"`
    CarYear string `json:"car_year" bson:"car_year"`
    CarMilage string `json:"car_milage" bson:"car_milage"`
    CarAddress string `json:"car_address" bson:"car_address"`
    CarName string `json:"car_name" bson:"car_name"`
    CarValue string `json:"car_value" bson:"car_value"`
    CarInsuranceInfo CarInsuranceInfo `json:"car_insurance_info" bson:"car_insurance_info"`
    CarFetures CarFetures `json:"car_fetures" bson:"car_fetures"`    
}
type CarFetures struct {
    WindowType string `json:"window_type" bson:"window_type"` // powerWindow , etc
    SoundType string `json:"sound_type" bson:"sound_type"` 
    LockingType string `json:"locking_type" json:"locking_type"`
    ParkingCensor string `json:"parking_censor" bson:"parking_censor"`
    AirBags bool `json:"air_bags" bson:"air_bags"`
    GearType string `json:"gear_type" bson:"gear_type"` // automatic,manual
    Camera bool `json:"camera" bson:"camera"` // ivrm
    DoorLock string `json:"door_lock" bson:"door_lock"`
    CarAc bool `json:"car_ac" bson:"car_ac"`
    PowerStreeing string `json:"power_steering" bson:"power_steering"`    
}
type CarInsuranceInfo struct {
    InsuranceCompanyName string `json:"insurance_company_name" bson:"insurance_company_name"`
    InsuranceProvider string `json:"insurance_provider" bson:"insurance_provider"`
    PolicyNumber string `json:"policy_number" bson:"policy_numer"`
    Website string `json:"website" bson:"website"`
}

// Customer Modules
type Customer struct {
    Id string `json:"_id" bson:"_id" `
    MobileNumber int64 `json:"mobile_number" bson:"mobile_number" `
    Email string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
}
type CustomerDetails struct {
    Id string `json:"_id" bson:"_id" `
    MobileNumber string `json:"mobile_number" bson:"mobile_number"`
    Name string `json:"namr" bson:"name"`
    CustomerAddress string `json:"customer_address" bson:"customer_address"`
    PassportNumber string `json:"passport_number" bson:"passport_number"`
    IDNumber string `json:"id_number" bson:"id_number"`
    DrvingLicenceNumber string `json:"driving_licence_number" bson:"driving_licence_number"`
    CurrentLocationLatitude string  `json:"current_location_latitude" bson:"current_location_latitude"`
    CurrentLocationLogtitude string  `json:"current_location_logtitude" bson:"current_location_logtitude"`  
}