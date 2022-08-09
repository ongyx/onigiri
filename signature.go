package onigiri

// Signature is a bitset with a fixed size of 64.
type Signature uint8

func (s *Signature) Test(n uint8) bool {
	return (*s & (1 << n)) == 1
}

func (s *Signature) Set(n uint8) {
	*s |= (1 << n)
}

func (s *Signature) Clear(n uint8) {
	*s &^= (1 << n)
}

func (s *Signature) Contains(sig Signature) bool {
	return (*s & sig) == sig
}
