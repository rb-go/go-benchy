package jwt_test

import (
	"fmt"
	"testing"
)

func BenchmarkJWT_Create_HS256_String_Empty(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateStringHS256(JWTID, Issuer, Subject, Audience, TimeToLive, plEmptyData)
				if err != nil {
					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_HS256_Bytes_Empty(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateBytesHS256(JWTID, Issuer, Subject, Audience, TimeToLive, plEmptyData)
				if err != nil {
					bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_HS256_String_Filled(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateStringHS256(JWTID, Issuer, Subject, Audience, TimeToLive, plData)
				if err != nil {
					bs.Fatalf("%s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_HS256_Bytes_Filled(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			bs.ResetTimer()
			bs.ReportAllocs()
			for i := 0; i < bs.N; i++ {
				_, err := stublist[isl].processor.CreateBytesHS256(JWTID, Issuer, Subject, Audience, TimeToLive, plData)
				if err != nil {
					bs.Fatalf("%s, err: %+v", taskKey, err)
				}
			}
		})
	}
	splitTests()
}
