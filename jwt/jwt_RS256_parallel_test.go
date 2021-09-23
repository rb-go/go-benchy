package jwt_test

import (
	"fmt"
	"testing"
)

func BenchmarkJWT_Create_RS256_String_Empty_Parallel(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			bs.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := stublist[isl].processor.CreateStringRS256(JWTID, Issuer, Subject, Audience, TimeToLive, plEmptyData)
					if err != nil {
						bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
					}
				}
			})
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			bs.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := stublist[isl].processor.CreateBytesRS256(JWTID, Issuer, Subject, Audience, TimeToLive, plEmptyData)
					if err != nil {
						bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
					}
				}
			})
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_RS256_String_Filled_Parallel(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			bs.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := stublist[isl].processor.CreateStringRS256(JWTID, Issuer, Subject, Audience, TimeToLive, plData)
					if err != nil {
						bs.Fatalf("%s, err: %+v", taskKey, err)
					}
				}
			})
		})
	}
	splitTests()
}

func BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			bs.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := stublist[isl].processor.CreateBytesRS256(JWTID, Issuer, Subject, Audience, TimeToLive, plData)
					if err != nil {
						b.Fatalf("%s, err: %+v", taskKey, err)
					}
				}
			})
		})
	}
	splitTests()
}
