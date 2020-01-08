package StructConfig

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Company struct {
	Id                        string `json:"_id" bson:"_id" validate:"required"`
	CreatedOn                 int64  `json:"created_on" bson:"created_on" validate:"required"`
	CompanyName               string `json:"company_name" bson:"company_name" validate:"required"`
	CompanyRegistrationNumber string `json:"company_registration_number" bson:"company_registration_number" validate:"required"`
	Email                     string `json:"email" bson:"email" validate:"required"`
	PhoneNumber               string `json:"phone_number" bson:"phone_number" validate:"required"`
	MobileNumber              string `json:"mobile_number" bson:"mobile_number" validate:"required"`
	UserId                    string `json:"user_id" bson:"user_id" validate:"required"`
	Password                  string `json:"password" bson:"password" validate:"required"`
	UserRole                  string `json:"user_role" bson:"user_role" validate:"required"`
}
type CompanyLocation struct {
	Id              string `json:"_id" bson:"_id"`
	MobileNumber    string `json:"mobile_number" bson:"mobile_number" validate:"required"`
	LocationAddedOn int64  `json:"location_added_on" bson:"location_added_on" validate:"required"`
	LocationName    string `json:"location_name" bson:"location_name" validate:"required"`
	CoodinatesX     string `json:"coordinates_x" bson:"coordinates_x" validate:"required"`
	CoordinatesY    string `json:"coordinates_y" bson:"coordinates_y" validate:"required"`
}
type UserInstance struct {
	MobileNumber string `json:"mobile_number" bson:"mobile_number" validate:"required"` // common in comapny and user
	Id           string `json:"_id" bson:"_id" validate:"required"`
	UserAddedOn  int64  `json:"user_added_on" bson:"user_added_on" validate:"required"`
	UserEmail    string `json:"user_email" bson:"user_email" validate:"required"`
	UserName     string `json:"user_name" bson:"user_name" validate:"required"`
	UserRole     string `json:"user_role" bson:"user_role" validate:"required"`
	Password     string `json:"password" bson:"password" validate:"required"`
}
type JwtStruct struct {
	UserEmail string `json:"user_email"`
	UserName  string `json:"user_name"`
	jwt.StandardClaims
}
type CarMaster struct {
	Id          string `json:"_id" bson:"_id" validate:"required"`
	AddedOn     int64  `json:"added_on" bson:"added_on" validate:"required"`
	CarCategory string `json:"car_category" bson:"car_category" validate:"required"`
	CarMake     string `json:"car_make" bson:"car_make" validate:"required"`
}
type CompanyCar struct {
	Id                string           `json:"_id" bson:"_id" validate:"required"`
	CompanyCarAddedOn int64            `json:"company_car_addedon" bson:"company_car_addedon" validate:"required"`
	CarCategoryId     string           `json:"car_category_id" bson:"car_category_id" validate:"required"`
	CarNumberPLate    string           `json:"car_number_plate" bson:"car_number_plate" validate:"required"`
	CarColor          string           `json:"car_color" bson:"car_color" validate:"required"`
	CarModel          string           `json:"car_model" bson:"car_model" validate:"required"`
	CarYear           string           `json:"car_year" bson:"car_year" validate:"required"`
	CarMilage         string           `json:"car_milage" bson:"car_milage" validate:"required"`
	CarAddress        string           `json:"car_address" bson:"car_address" validate:"required"`
	CarName           string           `json:"car_name" bson:"car_name" validate:"required"`
	CarValue          string           `json:"car_value" bson:"car_value" validate:"required"`
	CarInsuranceInfo  CarInsuranceInfo `json:"car_insurance_info" bson:"car_insurance_info" validate:"required"`
	CarFetures        CarFetures       `json:"car_fetures" bson:"car_fetures" validate:"required"`
}
type CarFetures struct {
	WindowType    string `json:"window_type" bson:"window_type" validate:"required"` // powerWindow , etc
	SoundType     string `json:"sound_type" bson:"sound_type" validate:"required"`
	LockingType   string `json:"locking_type" json:"locking_type" validate:"required"`
	ParkingCensor string `json:"parking_censor" bson:"parking_censor" validate:"required"`
	AirBags       bool   `json:"air_bags" bson:"air_bags" validate:"required"`
	GearType      string `json:"gear_type" bson:"gear_type" validate:"required"` // automatic,manual
	Camera        bool   `json:"camera" bson:"camera" validate:"required"`       // ivrm
	DoorLock      string `json:"door_lock" bson:"door_lock" validate:"required"`
	CarAc         bool   `json:"car_ac" bson:"car_ac" validate:"required"`
	PowerStreeing string `json:"power_steering" bson:"power_steering" validate:"required"`
}
type CarInsuranceInfo struct {
	InsuranceCompanyName string `json:"insurance_company_name" bson:"insurance_company_name" validate:"required"`
	InsuranceProvider    string `json:"insurance_provider" bson:"insurance_provider" validate:"required"`
	PolicyNumber         string `json:"policy_number" bson:"policy_numer" validate:"required"`
	Website              string `json:"website" bson:"website" validate:"required"`
}

// Customer Modules
type Customer struct {
	Id           string `json:"_id" bson:"_id" validate:"required"`
	MobileNumber int64  `json:"mobile_number" bson:"mobile_number" validate:"required" `
	Email        string `json:"email" bson:"email" validate:"required"`
	Password     string `json:"password" bson:"password" validate:"required"`
}
type CustomerDetails struct {
	Id                       string `json:"_id" bson:"_id" validate:"required" `
	MobileNumber             string `json:"mobile_number" bson:"mobile_number" validate:"required"`
	Name                     string `json:"namr" bson:"name" validate:"required"`
	CustomerAddress          string `json:"customer_address" bson:"customer_address" validate:"required"`
	PassportNumber           string `json:"passport_number" bson:"passport_number" validate:"required"`
	IDNumber                 string `json:"id_number" bson:"id_number" validate:"required"`
	DrvingLicenceNumber      string `json:"driving_licence_number" bson:"driving_licence_number" validate:"required"`
	CurrentLocationLatitude  string `json:"current_location_latitude" bson:"current_location_latitude" validate:"required"`
	CurrentLocationLogtitude string `json:"current_location_logtitude" bson:"current_location_logtitude" validate:"required"`
}
