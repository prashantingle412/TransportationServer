package TokenManager

import (
	jwt "github.com/dgrijalva/jwt-go"
	"TransportationServer/packages/StructConfig"
	// "TransportationServer/packages/RentalCompanyApi"	
	"log"
	// "net/http"
	// "fmt"
)

func GenerateToken(empStruct StructConfig.UserInstance)(string,error){
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &StructConfig.JwtStruct{
	  UserName : empStruct.UserName,
	  UserEmail : empStruct.UserEmail,
	})
	tokenstring, err := token.SignedString([]byte("write_some_secret_key_here"))
	if err != nil {
	  log.Println("Error while generate token : ",err)
	  return "",err
	} else {
	  return tokenstring,nil
	}
  }

  func DecodeToken(tokenString string)(interface{},error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	  return []byte("write_some_secret_key_here"), nil
	})
	if err != nil {
	  log.Println("Error while DecodeToken : ",err)
	  return "",err
	} else {
	  return token.Claims,nil
	}
	//return make(map[string]interface{}),nil
  }
  
  func IsTokenValid(tokenString string)(bool,error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	  return []byte("write_some_secret_key_here"), nil
	})
	// When using `Parse`, the result `Claims` would be a map.
	if err != nil {
	  log.Println("Error while IsTokenValid : ",err)
	  return false,err
	} else {
	  return token.Valid,nil
	}
  }
  