package utils

import (
	// redis "github.com/coroo/go-starter/config/redis"
	// jwt "github.com/dgrijalva/jwt-go"
	// "encoding/base64"
	"os"
	
	// entity "github.com/coroo/go-starter/app/entity"
	// dto "github.com/coroo/go-starter/app/dto"
	// "time"
)

func CreateAuth() string {
	// auth := os.Getenv("BASIC_AUTH_USERNAME") + ":" + os.Getenv("BASIC_AUTH_PASSWORD")
	// return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	return "uuid "+os.Getenv("AUTH_UUID");
}

const (
    tokenDuration = 72
    expireOffset  = 3600
)

// func CreateToken(res entity.User) (*dto.Token, error) {
// 	dtoToken := new(dto.Token)
// 	sign := jwt.New(jwt.GetSigningMethod("HS256"))
// 	// Set claims
// 	// This is the information which frontend can use
// 	// The backend can also decode the token and get admin etc.
// 	claims := sign.Claims.(jwt.MapClaims)
// 	claims["user_id"] = res.ID
// 	claims["email"] = res.Email
// 	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
// 	token, err := sign.SignedString([]byte("secret"))
// 	if err != nil {
// 		return dtoToken, err
// 	}

// 	refreshToken := jwt.New(jwt.SigningMethodHS256)
// 	rtClaims := refreshToken.Claims.(jwt.MapClaims)
// 	rtClaims["user_id"] = res.ID
// 	rtClaims["email"] = res.Email
// 	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
// 	rt, err := refreshToken.SignedString([]byte("secret"))
// 	if err != nil {
// 		return dtoToken, err
// 	}

// 	dtoToken.AccessToken = token
// 	dtoToken.RefreshToken  = rt
// 	return dtoToken, nil
// }

// func Logout(tokenString string, token *jwt.Token) error {
//     redisConn := redis.Connect()
// 	claims, ok := token.Claims.(jwt.MapClaims);
// 	if ok {
// 		return redisConn.SetValue(tokenString, tokenString, getTokenRemainingValidity(claims["exp"]))
// 	}
// 	return nil
// }

// func getTokenRemainingValidity(timestamp interface{}) int {
//     if validity, ok := timestamp.(float64); ok {
//         tm := time.Unix(int64(validity), 0)
//         remainer := tm.Sub(time.Now())
//         if remainer > 0 {
//             return int(remainer.Seconds() + expireOffset)
//         }
//     }
//     return expireOffset
// }

// func IsInBlacklist(token string) bool {
//     redisConn := redis.Connect()
//     redisToken, _ := redisConn.GetValue(token)
    
//     if redisToken == nil || redisToken == ""{
//         return false
//     }
    
//     return true
// }