package idgen_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	rs "github.com/riftbit/go-benchy/idgen/rs_xid"

	shaxbee "github.com/riftbit/go-benchy/idgen/shaxbee_go-snowflake"

	piaohua "github.com/riftbit/go-benchy/idgen/piaohua_snowflake-golang"
	sony "github.com/riftbit/go-benchy/idgen/sony_sonyflake"

	"github.com/riftbit/go-benchy/idgen/mongo-driver_mongo"

	"github.com/riftbit/go-benchy/idgen"
	bwmarrin "github.com/riftbit/go-benchy/idgen/bwmarrin_snowflake"
	godruoyi "github.com/riftbit/go-benchy/idgen/godruoyi_go-snowflake"
)

var bwmarrin_sf idgen.Snowflaker
var godruoyi_sf idgen.Snowflaker
var rs_sf idgen.Snowflaker
var shaxbee_sf idgen.Snowflaker
var sony_sf idgen.Snowflaker
var mongo_sf idgen.Snowflaker
var piaohua_sf idgen.Snowflaker

type stubs struct {
	name      string
	processor idgen.Snowflaker
	id        uint64
	hex       []byte
}

var stublist []stubs

func init() {

	// https://github.com/sony/sonyflake
	sony_sf = sony.NewSF(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	stublist = append(stublist, stubs{name: "sonyflake", processor: sony_sf})

	// https://github.com/shaxbee/go-snowflake
	shaxbee_sf = shaxbee.NewSF(1023)
	stublist = append(stublist, stubs{name: "shaxbee", processor: shaxbee_sf})

	// github.com/piaohua/snowflake-golang
	piaohua_sf = piaohua.NewSF(1)
	stublist = append(stublist, stubs{name: "piaohua", processor: piaohua_sf})

	// https://github.com/bwmarrin/snowflake
	bwmarrin_sf = bwmarrin.NewSF(1023)
	stublist = append(stublist, stubs{name: "bwmarrin", processor: bwmarrin_sf})

	// https://github.com/godruoyi/go-snowflake
	godruoyi_sf = godruoyi.NewSF()
	stublist = append(stublist, stubs{name: "godruoyi", processor: godruoyi_sf})

	// https://github.com/rs/xid
	rs_sf = rs.NewSF()
	stublist = append(stublist, stubs{name: "rs", processor: rs_sf})

	// https://go.mongodb.org/mongo-driver/mongo
	mongo_sf = mongo.NewSF()
	stublist = append(stublist, stubs{name: "mongo", processor: mongo_sf})
}

func splitTests() {
	fmt.Println(strings.Repeat("=", 155))
}

func BenchmarkSnowflake_CreateID(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateID()
				if err != nil {
					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

func BenchmarkSnowflake_CreateHex(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateHex()
				if err != nil {
					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

//
//func BenchmarkSnowflake_ParseID(b *testing.B) {
//	for isl := range stublist {
//		taskKey := fmt.Sprintf("%s", stublist[isl].name)
//		b.Run(taskKey, func(bs *testing.B) {
//			bs.ResetTimer()
//			bs.ReportAllocs()
//			for i := 0; i < bs.N; i++ {
//				err := stublist[isl].processor.ParseID(stublist[isl].id)
//				if err != nil {
//					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
//				}
//			}
//		})
//	}
//	splitTests()
//}
//
//func BenchmarkSnowflake_ParseHex(b *testing.B) {
//	for isl := range stublist {
//		taskKey := fmt.Sprintf("%s", stublist[isl].name)
//		b.Run(taskKey, func(bs *testing.B) {
//			bs.ResetTimer()
//			bs.ReportAllocs()
//			for i := 0; i < bs.N; i++ {
//				err := stublist[isl].processor.ParseHex(stublist[isl].hex)
//				if err != nil {
//					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
//				}
//			}
//		})
//	}
//	splitTests()
//}
