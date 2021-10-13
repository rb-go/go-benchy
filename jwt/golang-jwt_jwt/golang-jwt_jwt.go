package golang_jwt

import (
	"crypto/rsa"
	"time"

	jwtb "github.com/golang-jwt/jwt/v4"
	"github.com/riftbit/go-benchy/jwt"
)

// https://github.com/golang-jwt/jwt

type JWT struct {
	key        []byte
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWT(key []byte, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	return &JWT{
		key:        key,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

type MyCustomClaims struct {
	jwtb.RegisteredClaims
	jwt.PayloadData
}

func prepareToken(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) MyCustomClaims {
	return MyCustomClaims{
		RegisteredClaims: jwtb.RegisteredClaims{
			Audience:  aud,
			ID:        jti,
			Issuer:    iss,
			Subject:   sub,
			IssuedAt:  jwtb.NewNumericDate(time.Now()),
			ExpiresAt: jwtb.NewNumericDate(time.Now().Add(ttl)),
			NotBefore: jwtb.NewNumericDate(time.Now()),
		},
		PayloadData: payload,
	}
}

func (j *JWT) CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	// Sign and get the complete encoded token as a string using the secret
	return jwtb.NewWithClaims(jwtb.SigningMethodHS256, claims).SignedString(j.key)
}

func (j *JWT) CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	// Sign and get the complete encoded token as a string using the secret
	token, err := jwtb.NewWithClaims(jwtb.SigningMethodHS256, claims).SignedString(j.key)
	if err != nil {
		return nil, err
	}
	return []byte(token), nil
}

func (j *JWT) CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	// Sign and get the complete encoded token as a string using the secret
	return jwtb.NewWithClaims(jwtb.SigningMethodRS256, claims).SignedString(j.privateKey)
}

func (j *JWT) CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	// Sign and get the complete encoded token as a string using the secret
	token, err := jwtb.NewWithClaims(jwtb.SigningMethodRS256, claims).SignedString(j.privateKey)
	if err != nil {
		return nil, err
	}
	return []byte(token), nil
}

func (j *JWT) ValidateStrHS256(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytesHS256(token []byte) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateStrRS256(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytesRS256(token []byte) (bool, error) {
	return true, nil
}
