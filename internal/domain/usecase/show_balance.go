package usecase

import "github.com/dmitruk-v/piggy-bank/internal/domain/entity"

type ShowBalanceUseCaseInput interface {
	Execute() error
}

type ShowBalanceUseCase struct {
	balance *entity.Balance
}

func NewShowBalanceUseCase(balance *entity.Balance) *ShowBalanceUseCase {
	return &ShowBalanceUseCase{
		balance: balance,
	}
}

func (ucase *ShowBalanceUseCase) Execute() error {
	return nil
}
