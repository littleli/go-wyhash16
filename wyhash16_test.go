package wyhash16

import "testing"

var vec0 = [...]uint16{
	36519,
	6808,
	42654,
	12957,
	48725,
	19014,
	54864,
	25199,
	61043,
	31236,
	1653,
}

var vec11 = [...]uint16{
	43968,
	14321,
	50149,
	20468,
	56074,
	26607,
	62443,
	32710,
	3029,
	38877,
	9118,
	44928,
	15255,
	51107,
	21452,
	57290,
	26665,
	62521,
	32826,
	3135,
}

func TestHash0(t *testing.T) {
	h0 := NewHash(0)
	for n, want := range vec0 {
		got := h0.Next()
		if got != want {
			t.Errorf("hash.Next()[%d]=%d, want %d", n, got, want)
		}
	}
}

func TestHash11(t *testing.T) {
	h11 := NewHash(11)
	for n, want := range vec11 {
		got := h11.Next()
		if got != want {
			t.Errorf("hash.Next()[%d]=%d, want %d", n, got, want)
		}
	}
}
