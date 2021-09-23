package brianvoe

import (
	"time"

	sjwt "github.com/brianvoe/sjwt"
	"github.com/rb-pkg/benchy/jwt"
)

// https://github.com/brianvoe/sjwt
// Унылое говно!!! не поддерживает другие типы шифрования, работает пздц медленно,
// работает только со строками в качестве токена

type JWT struct {
	key []byte
}

func NewJWT(key []byte) *JWT {
	return &JWT{key: key}
}

func (j *JWT) CreateString(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {

	// Set Claims
	claims, err := sjwt.ToClaims(payload)
	if err != nil {
		return "", err
	}
	claims.Set(sjwt.TokenID, jti)            // UUID generated
	claims.SetIssuer(iss)                    // Issuer of the token
	claims.SetSubject(sub)                   // Subject of the token
	claims.SetIssuedAt(time.Now())           // IssuedAt in time, value is set in unix
	claims.SetNotBeforeAt(time.Now())        // Token valid in 1 hour
	claims.SetExpiresAt(time.Now().Add(ttl)) // Token expires in 24 hours
	claims.SetAudience(aud)                  // Audience the toke is for

	// Generate jwt
	return claims.Generate(j.key), nil
}

func (j *JWT) CreateBytes(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {

	// Set Claims
	claims, err := sjwt.ToClaims(payload)
	if err != nil {
		return nil, err
	}
	claims.Set(sjwt.TokenID, jti)            // UUID generated
	claims.SetIssuer(iss)                    // Issuer of the token
	claims.SetSubject(sub)                   // Subject of the token
	claims.SetIssuedAt(time.Now())           // IssuedAt in time, value is set in unix
	claims.SetNotBeforeAt(time.Now())        // Token valid in 1 hour
	claims.SetExpiresAt(time.Now().Add(ttl)) // Token expires in 24 hours
	claims.SetAudience(aud)                  // Audience the toke is for

	// Generate jwt
	return []byte(claims.Generate(j.key)), nil
}

func (j *JWT) ValidateStr(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytes(token []byte) (bool, error) {
	return true, nil
}
