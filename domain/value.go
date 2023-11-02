package domain

import "encoding/json"

type ValueObject[T any] struct {
	value T
}

func (v ValueObject[T]) Value() T {
	return v.value
}

func (v *ValueObject[T]) MarshalJSON() ([]byte, error) {
	type pubValueObject[T any] struct {
		Value T `json:"value"`
	}
	return json.Marshal(&pubValueObject[T]{Value: v.value})
}
