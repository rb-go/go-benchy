package golang_jwt

import (
	"time"

	jwtb "github.com/golang-jwt/jwt/v4"
	"github.com/rb-pkg/benchy/jwt"
)

// https://github.com/golang-jwt/jwt

type JWT struct {
	key []byte
}

func NewJWT(key []byte) *JWT {
	return &JWT{key: key}
}

type MyCustomClaims struct {
	jwtb.RegisteredClaims
	jwt.PayloadData
}

func (j *JWT) CreateString(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := MyCustomClaims{
		RegisteredClaims: jwtb.RegisteredClaims{
			Audience:  aud,
			ExpiresAt: jwtb.NewNumericDate(time.Now().Add(ttl)),
			ID:        jti,
			IssuedAt:  jwtb.NewNumericDate(time.Now()),
			Issuer:    iss,
			NotBefore: jwtb.NewNumericDate(time.Now()),
			Subject:   sub,
		},
		PayloadData: payload,
	}
	// Sign and get the complete encoded token as a string using the secret
	return jwtb.NewWithClaims(jwtb.SigningMethodHS256, claims).SignedString(j.key)
}

func (j *JWT) CreateBytes(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := MyCustomClaims{
		RegisteredClaims: jwtb.RegisteredClaims{
			Audience:  aud,
			ExpiresAt: jwtb.NewNumericDate(time.Now().Add(ttl)),
			ID:        jti,
			IssuedAt:  jwtb.NewNumericDate(time.Now()),
			Issuer:    iss,
			NotBefore: jwtb.NewNumericDate(time.Now()),
			Subject:   sub,
		},
		PayloadData: payload,
	}
	// Sign and get the complete encoded token as a string using the secret
	token, err := jwtb.NewWithClaims(jwtb.SigningMethodHS256, claims).SignedString(j.key)
	if err != nil {
		return nil, err
	}
	return []byte(token), nil
}

func (j *JWT) ValidateStr(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytes(token []byte) (bool, error) {
	return true, nil
}
