package domain

//go:generate mockgen -source=$GOFILE -destination=../mocks/domain/mock_$GOFILE -package=mock_domain

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
