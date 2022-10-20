package usecase

import "github.com/dmitruk-v/piggy-bank/internal/domain/entity"

type ShowBalanceUseCaseInput interface {
	Execute() error
}

type ShowBalanceUseCaseOutput interface {
	Present(res ShowBalanceResponse) error
}

type ShowBalanceResponse struct {
	Balance *entity.Balance
}

type ShowBalanceUseCase struct {
	balance *entity.Balance
	output  ShowBalanceUseCaseOutput
}

func NewShowBalanceUseCase(balance *entity.Balance, output ShowBalanceUseCaseOutput) *ShowBalanceUseCase {
	return &ShowBalanceUseCase{
		balance: balance,
		output:  output,
	}
}

func (ucase *ShowBalanceUseCase) Execute() error {
	return ucase.output.Present(ShowBalanceResponse{
		Balance: ucase.balance,
	})
}
