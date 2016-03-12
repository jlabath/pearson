// Package pearson implements Pearson hash function.
// See
// https://en.wikipedia.org/wiki/Pearson_hashing
package pearson

import "hash"

type ph struct {
	size    int
	results []uint8
	write   func([]byte) (int, error)
}

func (h *ph) BlockSize() int {
	return 1
}

func (h *ph) Reset() {
	for i := 0; i < h.size; i++ {
		h.results[i] = 0
	}
	h.write = h.firstWrite
}

func (h *ph) Size() int {
	return h.size
}

func (h *ph) Sum(p []byte) []byte {
	for i := 0; i < h.size; i++ {
		p = append(p, byte(h.results[i]))
	}
	return p
}

func (h *ph) Write(data []byte) (int, error) {
	return h.write(data)
}

func (h *ph) firstWrite(data []byte) (int, error) {
	for i := 0; i < h.size; i++ {
		h.results[i] = table[(int(data[0])+i)%256]
		for _, b := range data[1:] {
			h.results[i] = h.results[i] ^ table[uint8(b)]
		}
	}
	h.write = h.nextWrite
	return len(data), nil
}

func (h *ph) nextWrite(data []byte) (int, error) {
	for i := 0; i < h.size; i++ {
		for _, b := range data {
			h.results[i] = h.results[i] ^ table[uint8(b)]
		}
	}
	return len(data), nil
}

func newPh(size int) hash.Hash {
	res := make([]uint8, size)
	h := &ph{
		size:    size,
		results: res,
	}
	h.write = h.firstWrite
	return h
}

//New returns a new 8-bit Pearson hash.Hash.
func New() hash.Hash {
	return newPh(1)
}

//New16 returns a new 16-bit Pearson hash.Hash.
func New16() hash.Hash {
	return newPh(2)
}

//New24 returns a new 24-bit Pearson hash.Hash.
func New24() hash.Hash {
	return newPh(3)
}
