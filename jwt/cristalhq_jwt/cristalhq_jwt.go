package cristalhq

import (
	"time"

	jwtb "github.com/cristalhq/jwt/v3"
	"github.com/rb-pkg/benchy/jwt"
)

// https://github.com/cristalhq/jwt

type JWT struct {
	key     []byte
	builder *jwtb.Builder
	signer  jwtb.Signer
}

type UserClaims struct {
	jwtb.RegisteredClaims
	jwt.PayloadData
}

func NewJWT(key []byte) *JWT {
	signer, _ := jwtb.NewSignerHS(jwtb.HS256, key)
	res := &JWT{
		key:     key,
		signer:  signer,
		builder: jwtb.NewBuilder(signer),
	}
	return res
}

func (j *JWT) CreateString(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwtb.RegisteredClaims{
			ID:        jti,
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			ExpiresAt: jwtb.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwtb.NewNumericDate(time.Now()),
			NotBefore: jwtb.NewNumericDate(time.Now()),
		},
		PayloadData: payload,
	}
	token, err := j.builder.Build(claims)
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func (j *JWT) CreateBytes(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := UserClaims{
		RegisteredClaims: jwtb.RegisteredClaims{
			ID:        jti,
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			ExpiresAt: jwtb.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwtb.NewNumericDate(time.Now()),
			NotBefore: jwtb.NewNumericDate(time.Now()),
		},
		PayloadData: payload,
	}
	token, err := j.builder.Build(claims)
	if err != nil {
		return nil, err
	}
	return token.Raw(), nil
}

func (j *JWT) ValidateStr(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytes(token []byte) (bool, error) {
	return true, nil
}
