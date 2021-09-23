package gbrlsnchs

import (
	"time"

	jwtb "github.com/gbrlsnchs/jwt/v3"
	"github.com/rb-pkg/benchy/jwt"
)

// https://github.com/gbrlsnchs/jwt
// очень неплохой пакет, не сильно медленнее чем cristalhq_jwt, разница в аллокациях не большая, нужно смотреть дальше на удобство и функционал

type JWT struct {
	key *jwtb.HMACSHA
}

type CustomPayload struct {
	jwtb.Payload
	jwt.PayloadData
}

func NewJWT(key []byte) *JWT {
	return &JWT{key: jwtb.NewHS256(key)}
}

func (j *JWT) CreateString(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	pl := CustomPayload{
		Payload: jwtb.Payload{
			Issuer:         iss,
			Subject:        sub,
			Audience:       aud,
			ExpirationTime: jwtb.NumericDate(time.Now().Add(ttl)),
			NotBefore:      jwtb.NumericDate(time.Now()),
			IssuedAt:       jwtb.NumericDate(time.Now()),
			JWTID:          jti,
		},
		PayloadData: payload,
	}
	token, err := jwtb.Sign(pl, j.key)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytes(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	pl := CustomPayload{
		Payload: jwtb.Payload{
			Issuer:         iss,
			Subject:        sub,
			Audience:       aud,
			ExpirationTime: jwtb.NumericDate(time.Now().Add(ttl)),
			NotBefore:      jwtb.NumericDate(time.Now()),
			IssuedAt:       jwtb.NumericDate(time.Now()),
			JWTID:          jti,
		},
		PayloadData: payload,
	}
	return jwtb.Sign(pl, j.key)
}

func (j *JWT) ValidateStr(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytes(token []byte) (bool, error) {
	return true, nil
}
