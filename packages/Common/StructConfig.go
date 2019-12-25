package Common
type Company struct {
    Id string `json:"_id" bson:"_id"`
	CreatedOn int64 `json:"created_on" bson:"created_on"`
	CompanyName string `json:"company_name" bson:"company_name"`
    CompanyRegistrationNumber string `json:"company_registration_number" bson:"company_registration_number"`
    CompanyEmail string `json:"company_email" bson:"company_email"`
    PhoneNumber string `json:"phone_number" bson:"phone_number"`
    MobileNumber string `json:"mobile_number" bson:"mobile_number"`
    UserId string       `json:"user_id" bson:"user_id"`
    Password string `json:"password" bson:"password"`
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
}
type JwtStruct struct {
    
}