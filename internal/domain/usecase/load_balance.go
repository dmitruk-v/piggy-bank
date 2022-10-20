package usecase

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

type LoadBalanceUseCaseInput interface {
	Execute() error
}

type LoadBalanceUseCase struct {
	balance   *entity.Balance
	opStorage entity.OperationStorage
}

func NewLoadBalanceUseCase(balance *entity.Balance, opStorage entity.OperationStorage) *LoadBalanceUseCase {
	return &LoadBalanceUseCase{
		balance:   balance,
		opStorage: opStorage,
	}
}

func (ucase *LoadBalanceUseCase) Execute() error {
	ops, err := ucase.opStorage.GetAll()
	if err != nil {
		return fmt.Errorf("create balance from operations: %v", err)
	}
	var opErr error
	for _, op := range ops {
		switch op.Optype {
		case entity.DepositOperation:
			opErr = ucase.balance.Add(op.Currency, op.Amount)
		case entity.WithdrawOperation:
			opErr = ucase.balance.Sub(op.Currency, op.Amount)
		default:
			opErr = fmt.Errorf("unknown operation type: %v", op.Optype)
		}
		if opErr != nil {
			return fmt.Errorf("create balance from operations: %v", opErr)
		}
	}
	return nil
}
