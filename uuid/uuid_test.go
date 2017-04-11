package uuid

import (
	"os/exec"
	"testing"
)

func TestUUID(t *testing.T) {
	list := make(map[string]struct{})
	n := 1 << 20

	for i := 0; i < n; i++ {
		id := New()
		_, ok := list[id]
		if ok {
			t.Fail()
		}

		if i == 0 {
			t.Log(id)
		}
	}

	t.Log(n, "unique UUID")
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := exec.Command("uuidgen").Output()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkNewFromSource(b *testing.B) {
	src := NewSource()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src.New()
	}
}
