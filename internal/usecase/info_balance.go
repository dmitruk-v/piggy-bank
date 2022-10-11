package usecase

import "github.com/dmitruk-v/piggy-bank/internal/domain"

type InfoBalanceUseCase struct {
	balance *domain.Balance
}

func NewInfoBalanceUseCase(balance *domain.Balance) *InfoBalanceUseCase {
	return &InfoBalanceUseCase{
		balance: balance,
	}
}

func (ucase *InfoBalanceUseCase) Execute() error {
	return nil
}
