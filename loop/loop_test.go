package loop

import "testing"

func BenchmarkLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		loop(15000000)
	}
}
