package commons

import "github.com/brianvoe/gofakeit/v7"

func NewUUID() TUUID {
	return TUUID(gofakeit.UUID())
}
