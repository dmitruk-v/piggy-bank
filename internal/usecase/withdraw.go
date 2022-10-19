package usecase

import (
	"fmt"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type WithdrawRequest struct {
	Currency domain.Currency
	Amount   float64
}

type WithdrawUseCaseInput interface {
	Execute(req WithdrawRequest) error
}

type WithdrawUseCase struct {
	balance   *domain.Balance
	opStorage domain.OperationStorage
}

func NewWithdrawUseCase(balance *domain.Balance, opStorage domain.OperationStorage) *WithdrawUseCase {
	return &WithdrawUseCase{
		balance:   balance,
		opStorage: opStorage,
	}
}

func (ucase *WithdrawUseCase) Execute(req WithdrawRequest) error {
	if err := ucase.balance.Sub(req.Currency, req.Amount); err != nil {
		return fmt.Errorf("execute withdraw operation: %v", err)
	}
	op := domain.NewCurrencyOperation(domain.WithdrawOperation, req.Currency, req.Amount, time.Now().Unix(), nil, nil)
	if err := ucase.opStorage.Save(op); err != nil {
		return fmt.Errorf("execute withdraw operation: %v", err)
	}
	return nil
}
