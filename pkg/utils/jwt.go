package utils

import (
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Claims struct {
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	UserType string `json:"user_type,omitempty"`
	jwt.StandardClaims
}

func GenerateToken(ID int, user_name string, user_type string) (string, error) {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Setting Setup, fail to parse 'config.json' : %v", err)
	}

	var secret = viper.GetString(`jwt_secret`)
	expired_time := viper.GetInt(`expire_jwt`)
	issuer := viper.GetString(`app.issuer`)
	var jwt_secret = []byte(secret)

	claims := &Claims{
		UserID:   strconv.Itoa(ID),
		UserName: user_name,
		UserType: user_type,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expired_time)).Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwt_secret)
}

func ParseToken(token string) (*Claims, error) {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Setting Setup, fail to parse 'config.json' : %v", err)
	}

	var secret = viper.GetString(`jwt_secret`)
	var jwt_secret = []byte(secret)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwt_secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
