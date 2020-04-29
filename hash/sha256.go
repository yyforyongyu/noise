package hash

import (
	"crypto/sha256"
	"hash"
)

type hashSha256 struct {
	name string
	h    hash.Hash
}

// BlockLen for sha256 should be 64.
func (s *hashSha256) BlockLen() int {
	return s.h.BlockSize()
}

func (s *hashSha256) New() hash.Hash {
	return sha256.New()
}

// HashLen for sha256 should be 32.
func (s *hashSha256) HashLen() int {
	return s.h.Size()
}

func (s *hashSha256) String() string {
	return s.name
}

func (s *hashSha256) Reset() {
	s.h.Reset()
}

func newSha256() Hash {
	return &hashSha256{
		name: "SHA256",
		h:    sha256.New(),
	}
}

func init() {
	Register("SHA256", newSha256)
}
