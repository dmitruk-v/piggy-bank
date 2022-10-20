package usecase

import (
	"fmt"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type WithdrawRequest struct {
	Currency entity.Currency
	Amount   float64
}

type WithdrawUseCaseInput interface {
	Execute(req WithdrawRequest) error
}

type WithdrawUseCase struct {
	balance   entity.Balance
	opStorage entity.OperationStorage
}

func NewWithdrawUseCase(balance entity.Balance, opStorage entity.OperationStorage) *WithdrawUseCase {
	return &WithdrawUseCase{
		balance:   balance,
		opStorage: opStorage,
	}
}

func (ucase *WithdrawUseCase) Execute(req WithdrawRequest) error {
	if err := ucase.balance.Sub(req.Currency, req.Amount); err != nil {
		return fmt.Errorf("execute withdraw operation: %v", err)
	}
	op := entity.NewCurrencyOperation(entity.WithdrawOperation, req.Currency, req.Amount, time.Now().Unix(), nil, nil)
	if err := ucase.opStorage.Save(op); err != nil {
		return fmt.Errorf("execute withdraw operation: %v", err)
	}
	return nil
}
