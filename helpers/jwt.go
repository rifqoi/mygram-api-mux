package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rifqoi/mygram-api-mux/config"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_KEY = []byte(config.GetEnv("JWT_SIGNING_KEY"))

type JWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Id    int    `json:"id"`
}

func GenerateToken(email string, id int) (string, error) {
	LOGIN_EXP_DURATION := time.Now().Add(time.Hour * 7 * 24)

	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: LOGIN_EXP_DURATION.Unix(),
		},
		Email: email,
		Id:    id,
	}
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	signedToken, err := token.SignedString(JWT_KEY)

	return signedToken, err
}
func ValidateToken(encToken string) (jwt.MapClaims, error) {
	// Cek signing method dari token
	token, err := jwt.Parse(encToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method.")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Invalid signing method.")
		}
		return JWT_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// Convert dari claims ke int64 agar bisa diubah ke tipe data time.Time
	exp := int64(claims["exp"].(float64))

	// Convert unix timestamp ke tipe data time.Time
	expiredAt := time.Unix(exp, 0)
	if time.Now().After(expiredAt) {
		return nil, fmt.Errorf("Token expired.")
	}

	return claims, err
}
