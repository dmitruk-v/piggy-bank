package entity

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	UAH Currency = "UAH"
	RUB Currency = "RUB"
)

// TODO: Best way to store money in integer format.
// For example: user input is 125.65, so we parse to float
// and multiply by 100, then keep integer 12565 in storage.
// On presentation side wi just divide that number by 100
// and got 125.65

type Currency2 struct {
	Code        string
	Description string
}

var hm = map[string]Currency2{
	"usd": {Code: "USD", Description: "Dollar USA"},
	"eur": {Code: "EUR", Description: "European euro"},
	"uah": {Code: "UAH", Description: "Ukrainian hrivnya"},
	"rub": {Code: "RUB", Description: "Russian rubl"},
}
