package onigiri

type Signature uint64

func (s *Signature) Test(n uint64) bool {
	return (*s & (1 << n)) == 1
}

func (s *Signature) Set(n uint64) {
	*s |= (1 << n)
}

func (s *Signature) Clear(n uint64) {
	*s &^= (1 << n)
}

func (s *Signature) Contains(sig Signature) bool {
	return (*s & sig) == sig
}
