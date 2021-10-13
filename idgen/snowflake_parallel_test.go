package idgen_test

import (
	"fmt"
	"testing"
)

func BenchmarkSnowflake_CreateID_Parallel(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			bs.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := stublist[isl].processor.CreateID()
					if err != nil {
						bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
					}
				}
			})
		})
	}
	splitTests()
}

func BenchmarkSnowflake_CreateHex_Parallel(b *testing.B) {
	for isl := range stublist {
		taskKey := fmt.Sprintf("%s", stublist[isl].name)
		b.Run(taskKey, func(bs *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			bs.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := stublist[isl].processor.CreateHex()
					if err != nil {
						bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
					}
				}
			})
		})
	}
	splitTests()
}

//
//func BenchmarkSnowflake_ParseID_Parallel(b *testing.B) {
//	for isl := range stublist {
//		taskKey := fmt.Sprintf("%s", stublist[isl].name)
//		b.Run(taskKey, func(bs *testing.B) {
//			b.ResetTimer()
//			b.ReportAllocs()
//			bs.RunParallel(func(pb *testing.PB) {
//				for pb.Next() {
//					err := stublist[isl].processor.ParseID(stublist[isl].id)
//					if err != nil {
//						bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
//					}
//				}
//			})
//		})
//	}
//	splitTests()
//}
//
//func BenchmarkSnowflake_ParseHex_Parallel(b *testing.B) {
//	for isl := range stublist {
//		taskKey := fmt.Sprintf("%s", stublist[isl].name)
//		b.Run(taskKey, func(bs *testing.B) {
//			b.ResetTimer()
//			b.ReportAllocs()
//			bs.RunParallel(func(pb *testing.PB) {
//				for pb.Next() {
//					err := stublist[isl].processor.ParseHex(stublist[isl].hex)
//					if err != nil {
//						bs.Fatalf("invalid jwt create: %s, err: %+v", taskKey, err)
//					}
//				}
//			})
//		})
//	}
//	splitTests()
//}
