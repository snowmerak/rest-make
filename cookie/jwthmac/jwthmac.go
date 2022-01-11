package jwthmac

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func Deserialize(secret, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token is not valid")
	}
}

func Serialize(secret string, claims jwt.StandardClaims, additional jwt.MapClaims) (string, error) {
	if additional == nil {
		additional = make(jwt.MapClaims)
	}
	if claims.Audience != "" {
		additional["aud"] = claims.Audience
	}
	if claims.ExpiresAt != 0 {
		additional["exp"] = claims.ExpiresAt
	}
	if claims.Id != "" {
		additional["jti"] = claims.Id
	}
	if claims.IssuedAt != 0 {
		additional["iat"] = claims.IssuedAt
	}
	if claims.Issuer != "" {
		additional["iss"] = claims.Issuer
	}
	if claims.NotBefore != 0 {
		additional["nbf"] = claims.NotBefore
	}
	if claims.Subject != "" {
		additional["sub"] = claims.Subject
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
