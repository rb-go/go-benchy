package pascaldekloe

import (
	"crypto/rsa"
	"time"

	jwtb "github.com/pascaldekloe/jwt"
	"github.com/riftbit/go-benchy/jwt"
)

// https://github.com/pascaldekloe/jwt
// map[string]interface{} - жутко неудобно
// парсинг токена будет геморный! Ну его нафиг

type JWT struct {
	key        []byte
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type UserClaims struct {
	jwt.PayloadData
}

func NewJWT(key []byte, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	return &JWT{
		key:        key,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func prepareToken(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) jwtb.Claims {
	claims := jwtb.Claims{
		Registered: jwtb.Registered{
			Issuer:    iss,
			Subject:   sub,
			Audiences: aud,
			Expires:   jwtb.NewNumericTime(time.Now().Add(ttl)),
			NotBefore: jwtb.NewNumericTime(time.Now()),
			Issued:    jwtb.NewNumericTime(time.Now()),
			ID:        jti,
		},
	}
	claims.Set = make(map[string]interface{})
	claims.Set["login"] = payload.Login
	claims.Set["password"] = payload.Password
	claims.Set["is_admin"] = payload.IsAdmin
	claims.Set["roles"] = payload.Roles
	claims.Set["refresh_token"] = payload.RefreshToken
	claims.Set["additional_data"] = payload.AdditionalData
	claims.Set["omit_empty_additional_data_filled"] = payload.OmitEmptyAdditionalDataFilled
	claims.Set["omit_empty_additional_data_empty"] = payload.OmitEmptyAdditionalDataEmpty
	claims.Set["avatar"] = payload.Avatar
	return claims
}

func (j *JWT) CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := claims.HMACSign(jwtb.HS256, j.key)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	return claims.HMACSign(jwtb.HS256, j.key)
}

func (j *JWT) CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := claims.RSASign(jwtb.RS256, j.privateKey)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	return claims.RSASign(jwtb.RS256, j.privateKey)
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
