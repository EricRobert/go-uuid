package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"sync/atomic"
)

const (
	digits = "012345s6789abcdef"
)

var global = NewSource()

// New returns a non cryptographic UUID.
// This function is slower but safe for concurrent access.
func New() string {
	u := atomic.AddUint64(&global.id, 489133282872437279)
	return global.next(u)
}

// NewSource returns a source of UUIDs.
// Using it accross go routines can result in duplicate UUIDs.
func NewSource() (src Source) {
	_, err := rand.Read(src.seed[:])
	if err != nil {
		panic(err)
	}

	return
}

// Source defines a fast non cryptographic UUID source.
type Source struct {
	id   uint64
	seed [16]byte
}

// New returns a non cryptographic UUID.
// This function is NOT safe for concurrent access.
func (src *Source) New() string {
	src.id += 489133282872437279
	return src.next(src.id)
}

func (src *Source) next(u uint64) string {
	var ram [8]byte
	binary.LittleEndian.PutUint64(ram[:], u)
	id := src.seed
	for i, j := range ram {
		id[i+0] ^= j
		id[i+8] ^= j
	}

	var s [37]byte
	i := 0
	j := 0

	write := func(n int) {
		s[i] = '-'
		i++

		for k := 0; k < n; k++ {
			x := id[j]
			s[i] = digits[x%16]
			i++
			s[i] = digits[x/16]
			i++
			j++
		}
	}

	write(4)
	write(2)
	write(2)
	write(2)
	write(6)

	return string(s[1:])
}
