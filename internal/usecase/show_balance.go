package usecase

import "github.com/dmitruk-v/piggy-bank/internal/domain"

type ShowBalanceUseCaseInput interface {
	Execute() error
}

type ShowBalanceUseCase struct {
	balance *domain.Balance
}

func NewShowBalanceUseCase(balance *domain.Balance) *ShowBalanceUseCase {
	return &ShowBalanceUseCase{
		balance: balance,
	}
}

func (ucase *ShowBalanceUseCase) Execute() error {
	return nil
}
