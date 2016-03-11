package pearson

import (
	"encoding/hex"
	"testing"
)

type TabTest struct {
	in  uint8
	out uint8
}

var tabTests = []TabTest{
	{0, 98},
	{254, 138},
	{112, 39},
}

func TestBasic(t *testing.T) {
	h := New()
	for _, tc := range tabTests {
		h.Reset()
		h.Write([]byte{byte(tc.in)})
		s := h.Sum(nil)
		if uint8(s[0]) != tc.out {
			t.Errorf("got %d wanted %d", uint8(s[0]), tc.out)
		}
		if len(s) != h.Size() {
			t.Errorf("got %d wanted %d", len(s), h.Size())
		}
		if h.BlockSize() != 1 {
			t.Errorf("got %d wanted %d", h.BlockSize(), 1)
		}
	}
}

type DigTest struct {
	in    string
	hex   string
	hex16 string
	hex24 string
}

var digTable = []DigTest{
	{"quick brown fox", "92", "92be", "92beb7"},
	{"food", "e8", "e8da", "e8da23"},
	{"doof", "e8", "e81b", "e81b00"},
	{"foo", "f8", "f8ca", "f8ca33"},
}

func TestDigests(t *testing.T) {
	for _, tc := range digTable {
		h := New()
		h.Write([]byte(tc.in))
		s := hex.EncodeToString(h.Sum(nil))
		if s != tc.hex {
			t.Errorf("got %s but wanted %s", s, tc.hex)
		}
	}
}

func TestStreaming(t *testing.T) {
	for _, tc := range digTable {
		h := New()
		for _, bt := range []byte(tc.in) {
			h.Write([]byte{bt})
		}
		s := hex.EncodeToString(h.Sum(nil))
		if s != tc.hex {
			t.Errorf("got %s but wanted %s", s, tc.hex)
		}
	}
}

func TestDigests16Streamin(t *testing.T) {
	for _, tc := range digTable {
		h := New16()
		for _, bt := range []byte(tc.in) {
			h.Write([]byte{bt})
		}
		s := hex.EncodeToString(h.Sum(nil))
		if s != tc.hex16 {
			t.Errorf("got %s but wanted %s", s, tc.hex16)
		}
		if h.Size() != 2 {
			t.Errorf(
				"16bit hash must be two bytes long but got %d",
				h.Size())
		}

	}
}

func TestDigests24Streamin(t *testing.T) {
	h := New24()
	for _, tc := range digTable {
		h.Reset()
		for _, bt := range []byte(tc.in) {
			h.Write([]byte{bt})
		}
		s := hex.EncodeToString(h.Sum(nil))
		if s != tc.hex24 {
			t.Errorf("got %s but wanted %s", s, tc.hex24)
		}
		if h.Size() != 3 {
			t.Errorf(
				"24bit hash must be three bytes long but got %d",
				h.Size())
		}

	}
}
