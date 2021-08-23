//go:generate go test -benchmem -cpu 1,2,4,8 -bench=.

package str_concat

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func Benchmark_Pluses(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	var s string
	//b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = "lorem" + rndVal + "ipsum " + "with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example " + " or something else :)"
		_ = s
	}
}

func Benchmark_Pluses_Parallel(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	var s string
	//b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s = "lorem" + rndVal + "ipsum " + "with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example " + " or something else :)"
			_ = s
		}
	})
}

func Benchmark_PlusesFromEnv(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	envVal := os.Getenv("GOPATH")
	var s string
	//b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = "lorem" + rndVal + envVal + "ipsum " + "with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example " + " or something else :)"
		_ = s
	}
}

func Benchmark_PlusesFromEnv_Parallel(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	envVal := os.Getenv("GOPATH")
	var s string
	//b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s = "lorem" + rndVal + envVal + "ipsum " + "with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example " + " or something else :)"
			_ = s
		}
	})
}

func Benchmark_Fmt(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	var s string
	//b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf("%s%s%s%s%s", "lorem", rndVal, "ipsum ", "with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example ", " or something else :)")
		_ = s
	}
}

func Benchmark_Fmt_Parallel(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	var s string
	//b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s = fmt.Sprintf("%s%s%s%s%s", "lorem", rndVal, "ipsum ", "with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example ", " or something else :)")
			_ = s
		}
	})
}

func Benchmark_StringBuilder(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	//b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		//sb.Grow(2048)
		var sb strings.Builder
		_, _ = sb.WriteString("lorem")
		_, _ = sb.WriteString(rndVal)
		_, _ = sb.WriteString("ipsum ")
		_, _ = sb.WriteString("with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example ")
		_, _ = sb.WriteString(" or something else :)")
		_ = sb.String()
		sb.Reset()
	}
}

func Benchmark_StringBuilder_Parallel(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	//b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var sb strings.Builder
			//sb.Grow(2048)
			_, _ = sb.WriteString("lorem")
			_, _ = sb.WriteString(rndVal)
			_, _ = sb.WriteString("ipsum ")
			_, _ = sb.WriteString("with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example ")
			_, _ = sb.WriteString(" or something else :)")
			_ = sb.String()
			sb.Reset()
		}
	})
}

func Benchmark_ByteBufferPool(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	//b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bb := bytebufferpool.Get()
		_, _ = bb.WriteString("lorem")
		_, _ = bb.WriteString(rndVal)
		_, _ = bb.WriteString("ipsum ")
		_, _ = bb.WriteString("with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example ")
		_, _ = bb.WriteString(" or something else :)")
		_ = bb.String()
		bytebufferpool.Put(bb)
	}
}

func Benchmark_ByteBufferPool_Parallel(b *testing.B) {
	// run the function b.N times
	rndVal := strconv.Itoa(rand.Int() + rand.Int())
	//b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb := bytebufferpool.Get()
			_, _ = bb.WriteString("lorem")
			_, _ = bb.WriteString(rndVal)
			_, _ = bb.WriteString("ipsum ")
			_, _ = bb.WriteString("with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example with long string for example ")
			_, _ = bb.WriteString(" or something else :)")
			_ = bb.String()
			bytebufferpool.Put(bb)
		}
	})
}
