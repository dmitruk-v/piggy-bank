package entity

import (
	"fmt"
	"strings"
)

type Currency int

const (
	USD Currency = iota + 1
	EUR
	UAH
	RUB
)

var currencies = [...]Currency{
	0,
	USD,
	EUR,
	UAH,
	RUB,
}

var names = [...]string{
	"",
	"USD",
	"EUR",
	"UAH",
	"RUB",
}

func (c Currency) String() string {
	return names[c]
}

func CurrencyFromString(src string) (Currency, error) {
	for c, name := range names {
		if strings.ToUpper(src) == name {
			return Currency(c), nil
		}
	}
	return Currency(-1), fmt.Errorf("currency %q not found", src)
}

func CurrencyFromInt(src int) (Currency, error) {
	if src < 1 || src > len(currencies)-1 {
		return Currency(-1), fmt.Errorf("currency %q not found", src)
	}
	return Currency(src), nil
}

func Currencies() []Currency {
	return currencies[1:]
}

// func (c Currency) IsValid() bool {
// 	switch c {
// 	case USD, EUR, UAH, RUB:
// 		return true
// 	}
// 	return false
// }

// TODO: Best way to store money in integer format.
// For example: user input is 125.65, so we parse to float
// and multiply by 100, then keep integer 12565 in storage.
// On presentation side we just divide that number by 100
// and got 125.65

// type Currencies struct {
// 	items []Currency
// }
