package iteration

import "testing"

func Repeat(character string) string {
	var repeated string
	const repeatcount = 5
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
