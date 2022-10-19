package usecase

import (
	"fmt"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type LoadBalanceUseCase struct {
	balance   *domain.Balance
	opStorage domain.OperationStorage
}

func NewLoadBalanceUseCase(balance *domain.Balance, opStorage domain.OperationStorage) *LoadBalanceUseCase {
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
		case domain.DepositOperation:
			opErr = ucase.balance.Add(op.Currency, op.Amount)
		case domain.WithdrawOperation:
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
