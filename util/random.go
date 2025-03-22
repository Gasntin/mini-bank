package util

import (
	"github.com/brianvoe/gofakeit/v7"
)

func RandomString(length int) string {
	return gofakeit.LetterN(uint(length))
}

func RandomInt(min, max int) int {
	return gofakeit.Number(min, max)
}

func RandomOwner() string {
	return gofakeit.Name()
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCurrency() string {
	return gofakeit.CurrencyShort()
}
