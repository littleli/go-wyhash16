package wyhash16

// WyHash16 holds state of the rng and magic odd number
type WyHash16 struct {
	rng uint16
	odd uint16
}

// NewHash generates new safe instance
func NewHash(init uint16) *WyHash16 {
	return &WyHash16{rng: init, odd: 0xfc15}
}

// Next gives next 16bit pseudo-random number
func (h *WyHash16) Next() uint16 {
	h.rng += h.odd
	return uint16(hash16(uint32(h.rng), 0x2ab))
}

// SetOdd allow to set different magic odd number
func (h *WyHash16) SetOdd(newOdd uint16) {
	h.odd = newOdd
}

func hash16(input, key uint32) uint32 {
	z := input * key
	return ((z >> 16) ^ z) & 0xffff
}
