/*
Package utils provided lots of tool functions
*/
package utils

import (
	"doslism/office-chat/types"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// KEY is the key to generate tokens
var KEY = []byte("ZiwenWangIsSvenKing")

/*
GenerateNewID is to generate an unique user Id
*/
func GenerateNewID() string {
	var randomStr string = uuid.New().String()
	ts := time.Now().UnixNano()
	return fmt.Sprintf("%d-%s", ts, randomStr)
}

/*
GenerateNewRefreshToken is to generate an refreshToken token
*/
func GenerateNewRefreshToken(uid string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["tokenType"] = "refresh"
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	claims["refreshToken"] = ""
	tokenString, err := token.SignedString(KEY)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
GenerateNewAccessToken is to generate an accessToken token
*/
func GenerateNewAccessToken(uid string, refreshToken string) (string, error) {
	if uid == "" || refreshToken == "" {
		return "", errors.New("uid or refreshToken is not valid")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["tokenType"] = "access"
	claims["uid"] = uid
	claims["refreshToken"] = refreshToken
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(KEY)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

/*
DecodeToken is to decode the token to an object
*/
func DecodeToken(tokenString string) (types.TokenPayload, error) {
	if tokenString == "" {
		return types.TokenPayload{}, errors.New("No token available")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return KEY, nil
	})

	if err != nil {
		return types.TokenPayload{}, err
	}

	if token == nil {
		return types.TokenPayload{}, errors.New("No token available")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return types.TokenPayload{}, errors.New("Token error")
	}

	var res types.TokenPayload

	if claims["tokenType"] == "refresh" || claims["refeshToken"] == nil {
		res = types.TokenPayload{
			Uid: claims["uid"].(string),
			// RefeshToken: claims["refeshToken"].(string),
			Exp:       claims["exp"].(float64),
			TokenType: claims["tokenType"].(string),
		}
	} else {
		res = types.TokenPayload{
			Uid:         claims["uid"].(string),
			RefeshToken: claims["refeshToken"].(string),
			Exp:         claims["exp"].(float64),
			TokenType:   claims["tokenType"].(string),
		}
	}

	return res, nil

}
