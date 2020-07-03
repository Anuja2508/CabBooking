package token

import (
	"GoCab/model"
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcec"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var (
	accessKey string
)

// CreateMobileToken - creates mobile token
func CreateMobileToken(PhoneNumber string) (string, error) {

	claims := &model.UserClaims{
		PhoneNumber: PhoneNumber,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(accessKey))
	return ss, err
}

// CreateAccessKey :- For creating random key for genrating access token
func CreateAccessKey() (string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", err
	}
	accessKey = hex.EncodeToString(privKey.D.Bytes())
	return accessKey, nil

}

// GetPhoneNumberFromJWT :- get phone number of login user
func GetPhoneNumberFromJWT(context echo.Context) (string, error) {
	if context.Get("user") == nil {
		return "", errors.New("Invalid token")
	}
	token := context.Get("user").(*jwt.Token)

	if token != nil {
		claims := token.Claims.(*model.UserClaims)

		return claims.PhoneNumber, nil
	} else {
		return "", errors.New("Invalid token")
	}
}
