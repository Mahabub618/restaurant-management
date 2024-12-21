package helper

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

func GenerateAllTokens(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {

}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {

}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {

}
