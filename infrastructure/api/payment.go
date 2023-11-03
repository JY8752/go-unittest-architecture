package api

//go:generate mockgen -source=$GOFILE -destination=../../mocks/api/mock_$GOFILE -package=mock_api

import "errors"

type Payment interface {
	Buy(int) error
}

type payment struct{}

func NewPayment() Payment {
	return &payment{}
}

// dummy実装 本当は外部の決済APIとかを叩く想定
func (p *payment) Buy(price int) error {
	if price == 100 {
		return nil
	}
	return errors.New("invalid price")
}
