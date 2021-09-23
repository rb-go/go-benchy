package pascaldekloe

import (
	"time"

	jwtb "github.com/pascaldekloe/jwt"
	"github.com/rb-pkg/benchy/jwt"
)

// https://github.com/pascaldekloe/jwt
// map[string]interface{} - жутко неудобно
// парсинг токена будет геморный! Ну его нафиг

type JWT struct {
	key []byte
}

type UserClaims struct {
	jwt.PayloadData
}

func NewJWT(key []byte) *JWT {
	return &JWT{key: key}
}

func (j *JWT) CreateString(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {

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

	token, err := claims.HMACSign(jwtb.HS256, j.key)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytes(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {

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

	return claims.HMACSign(jwtb.HS256, j.key)
}

func (j *JWT) ValidateStr(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytes(token []byte) (bool, error) {
	return true, nil
}
