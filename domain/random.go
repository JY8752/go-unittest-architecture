package domain

import "time"

type SeedGenerator interface {
	New() int64
}

type seedGenerator struct{}

func NewSeedGenerator() SeedGenerator {
	return &seedGenerator{}
}

func (s *seedGenerator) New() int64 {
	return time.Now().UnixNano()
}
