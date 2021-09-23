package jwt_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/rb-pkg/benchy/jwt"
	brianvoe "github.com/rb-pkg/benchy/jwt/brianvoe_sjwt"
	cristalhq "github.com/rb-pkg/benchy/jwt/cristalhq_jwt"
	gbrlsnchs "github.com/rb-pkg/benchy/jwt/gbrlsnchs_jwt"
	golang_jwt "github.com/rb-pkg/benchy/jwt/golang-jwt_jwt"
	lestrrat_go "github.com/rb-pkg/benchy/jwt/lestrrat-go_jwx"
	pascaldekloe "github.com/rb-pkg/benchy/jwt/pascaldekloe_jwt"
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

func init() {

	// https://github.com/golang-jwt/jwt
	golang_jwt_jwt = golang_jwt.NewJWT(signKeyB)
	stublist = append(stublist, stubs{"golang-jwt_jwt", golang_jwt_jwt})

	// https://github.com/lestrrat-go/jwx
	lestrrat_go_jwx = lestrrat_go.NewJWT(signKeyB)
	stublist = append(stublist, stubs{"lestrrat-go_jwx", lestrrat_go_jwx})

	// https://github.com/gbrlsnchs/jwt
	gbrlsnchs_jwt = gbrlsnchs.NewJWT(signKeyB)
	stublist = append(stublist, stubs{"gbrlsnchs_jwt", gbrlsnchs_jwt})

	// https://github.com/pascaldekloe/jwt
	pascaldekloe_jwt = pascaldekloe.NewJWT(signKeyB)
	stublist = append(stublist, stubs{"pascaldekloe_jwt", pascaldekloe_jwt})

	// https://github.com/brianvoe/sjwt
	brianvoe_sjwt = brianvoe.NewJWT(signKeyB)
	stublist = append(stublist, stubs{"brianvoe_sjwt", brianvoe_sjwt})

	// https://github.com/cristalhq/jwt
	cristalhq_jwt = cristalhq.NewJWT(signKeyB)
	stublist = append(stublist, stubs{"cristalhq_jwt", cristalhq_jwt})
}

func splitTests() {
	fmt.Println(strings.Repeat("=", 150))
}

func BenchmarkJWT_Create_String_Empty(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateString(JWTID, Issuer, Subject, Audience, TimeToLive, plEmptyData)
				if err != nil {
					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_Bytes_Empty(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateBytes(JWTID, Issuer, Subject, Audience, TimeToLive, plEmptyData)
				if err != nil {
					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_String_Filled(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateString(JWTID, Issuer, Subject, Audience, TimeToLive, plData)
				if err != nil {
					bs.Fatalf("%s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_Bytes_Filled(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateBytes(JWTID, Issuer, Subject, Audience, TimeToLive, plData)
				if err != nil {
					bs.Fatalf("%s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}
