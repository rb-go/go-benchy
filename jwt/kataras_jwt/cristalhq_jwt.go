package cristalhq

import (
	"crypto/rsa"
	"time"

	jwtb "github.com/kataras/jwt"
	"github.com/rb-pkg/benchy/jwt"
)

// https://github.com/cristalhq/jwt

type JWT struct {
	key        []byte
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type UserClaims struct {
	jwtb.Claims
	jwt.PayloadData
}

func NewJWT(key []byte, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	res := &JWT{
		key:        key,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	return res
}

func prepareToken(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) UserClaims {
	return UserClaims{
		Claims: jwtb.Claims{
			ID:        jti,
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			Expiry:    time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
		PayloadData: payload,
	}
}

func (j *JWT) CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := jwtb.Sign(jwtb.HS256, j.key, claims)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	return jwtb.Sign(jwtb.HS256, j.key, claims)
}

func (j *JWT) CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := jwtb.Sign(jwtb.RS256, j.privateKey, claims)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	return jwtb.Sign(jwtb.RS256, j.privateKey, claims)
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
