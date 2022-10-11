package usecase

import (
	"fmt"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type WithdrawUseCase struct {
	balance   *domain.Balance
	opStorage OperationStorage
}

func NewWithdrawUseCase(balance *domain.Balance, opStorage OperationStorage) *WithdrawUseCase {
	return &WithdrawUseCase{
		balance:   balance,
		opStorage: opStorage,
	}
}

func (ucase *WithdrawUseCase) Execute(currency domain.Currency, amount float64) error {
	if err := ucase.balance.Sub(currency, amount); err != nil {
		return fmt.Errorf("execute withdraw operation: %v", err)
	}
	op := domain.NewCurrencyOperation(domain.WithdrawOperation, currency, amount, time.Now().Unix())
	if err := ucase.opStorage.Save(op); err != nil {
		return fmt.Errorf("execute withdraw operation: %v", err)
	}
	return nil
}
