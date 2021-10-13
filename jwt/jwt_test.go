package jwt_test

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	kataras "github.com/riftbit/go-benchy/jwt/kataras_jwt"

	"github.com/riftbit/go-benchy/jwt"
	brianvoe "github.com/riftbit/go-benchy/jwt/brianvoe_sjwt"
	cristalhq "github.com/riftbit/go-benchy/jwt/cristalhq_jwt"
	gbrlsnchs "github.com/riftbit/go-benchy/jwt/gbrlsnchs_jwt"
	golang_jwt "github.com/riftbit/go-benchy/jwt/golang-jwt_jwt"
	lestrrat_go "github.com/riftbit/go-benchy/jwt/lestrrat-go_jwx"
	pascaldekloe "github.com/riftbit/go-benchy/jwt/pascaldekloe_jwt"
)

const JWTID = "ergoz"
const Issuer = "[riftbit] ErgoZ"
const Subject = "ergoz"
const TimeToLive = time.Hour * 7

var Audience = []string{"https://golang.org", "https://jwt.io"}

var golang_jwt_jwt jwt.JWTer
var lestrrat_go_jwx jwt.JWTer
var gbrlsnchs_jwt jwt.JWTer
var pascaldekloe_jwt jwt.JWTer
var brianvoe_sjwt jwt.JWTer
var cristalhq_jwt jwt.JWTer
var kataras_jwt jwt.JWTer

var plData = jwt.PayloadData{
	Login:                         "ergozru@gmail.com",
	Password:                      "oh@my%awesome777p@ssw0rd",
	IsAdmin:                       true,
	Roles:                         []string{"admin", "user", "manager"},
	RefreshToken:                  []byte{},
	AdditionalData:                json.RawMessage(`{"glossary":{"title":"example glossary","GlossDiv":{"title":"S","GlossList":{"GlossEntry":{"ID":"SGML","SortAs":"SGML","GlossTerm":"Standard Generalized Markup Language","Acronym":"SGML","Abbrev":"ISO 8879:1986","GlossDef":{"para":"A meta-markup language, used to create markup languages such as DocBook.","GlossSeeAlso":["GML","XML"]},"GlossSee":"markup"}}}}}`),
	OmitEmptyAdditionalDataFilled: json.RawMessage(`{"glossary":{"title":"example glossary","GlossDiv":{"title":"S","GlossList":{"GlossEntry":{"ID":"SGML","SortAs":"SGML","GlossTerm":"Standard Generalized Markup Language","Acronym":"SGML","Abbrev":"ISO 8879:1986","GlossDef":{"para":"A meta-markup language, used to create markup languages such as DocBook.","GlossSeeAlso":["GML","XML"]},"GlossSee":"markup"}}}}}`),
	OmitEmptyAdditionalDataEmpty:  nil,
	Avatar:                        jwt.AvatarImage,
}

var plEmptyData = jwt.PayloadData{
	Login:                         "",
	Password:                      "",
	IsAdmin:                       false,
	Roles:                         nil,
	RefreshToken:                  nil,
	AdditionalData:                nil,
	OmitEmptyAdditionalDataFilled: nil,
	OmitEmptyAdditionalDataEmpty:  nil,
	Avatar:                        nil,
}

type stubs struct {
	name      string
	processor jwt.JWTer
}

var stublist []stubs
var signKeyB = []byte(`secret`)
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	publicKey := &privateKey.PublicKey

	// https://github.com/cristalhq/jwt
	cristalhq_jwt = cristalhq.NewJWT(signKeyB, privateKey, publicKey)
	stublist = append(stublist, stubs{"cristalhq_jwt", cristalhq_jwt})

	// https://github.com/kataras/jwt
	kataras_jwt = kataras.NewJWT(signKeyB, privateKey, publicKey)
	stublist = append(stublist, stubs{"kataras_jwt", kataras_jwt})

	// https://github.com/gbrlsnchs/jwt
	gbrlsnchs_jwt = gbrlsnchs.NewJWT(signKeyB, privateKey, publicKey)
	stublist = append(stublist, stubs{"gbrlsnchs_jwt", gbrlsnchs_jwt})

	// https://github.com/golang-jwt/jwt
	golang_jwt_jwt = golang_jwt.NewJWT(signKeyB, privateKey, publicKey)
	stublist = append(stublist, stubs{"golang-jwt_jwt", golang_jwt_jwt})

	// https://github.com/pascaldekloe/jwt
	pascaldekloe_jwt = pascaldekloe.NewJWT(signKeyB, privateKey, publicKey)
	stublist = append(stublist, stubs{"pascaldekloe_jwt", pascaldekloe_jwt})

	// https://github.com/lestrrat-go/jwx
	lestrrat_go_jwx = lestrrat_go.NewJWT(signKeyB, privateKey, publicKey)
	stublist = append(stublist, stubs{"lestrrat-go_jwx", lestrrat_go_jwx})

	// https://github.com/brianvoe/sjwt
	brianvoe_sjwt = brianvoe.NewJWT(signKeyB, privateKey, publicKey)
	stublist = append(stublist, stubs{"brianvoe_sjwt", brianvoe_sjwt})
}

func splitTests() {
	fmt.Println(strings.Repeat("=", 155))
}
