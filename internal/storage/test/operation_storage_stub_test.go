package test

import (
	"testing"

	"github.com/dmitruk-v/piggy-bank/internal/common"
	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
	"github.com/dmitruk-v/piggy-bank/internal/storage"
)

var (
	bcService common.BlockchainService
	opCreator entity.OperationCreator
)

func init() {
	bcService = common.NewBlockchainServiceImpl()
	opCreator = entity.NewOperationsCreatorImpl(bcService)
}

func TestOperationStorageStub_GetAll_Test(t *testing.T) {
	op1, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 12345, 1398351430, nil)
	op2, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 12345, 1398351431, op1.Hash)
	opStorage := storage.NewStubOperationStorage([]*entity.CurrencyOperation{op1, op2})
	ops, err := opStorage.GetAll()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if ops[0] != op1 {
		t.Errorf("got: %v, want: %v", ops[0], op1)
	}
	if ops[1] != op2 {
		t.Errorf("got: %v, want: %v", ops[1], op2)
	}
}

func TestOperationStorageStub_GetLatest_Test(t *testing.T) {
	op1, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 12345, 1398351430, nil)
	op2, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 12345, 1398351431, op1.Hash)
	opStorage := storage.NewStubOperationStorage([]*entity.CurrencyOperation{op1, op2})
	latest, err := opStorage.GetLatest()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if op2 != latest {
		t.Errorf("got: %v, want: %v", latest, op2)
	}
}

func TestOperationStorageStub_Save_Test(t *testing.T) {
	op1, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 12345, 1398351430, nil)
	op2, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 12345, 1398351431, op1.Hash)
	op3, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 12345, 1398351432, op2.Hash)
	opStorage := storage.NewStubOperationStorage([]*entity.CurrencyOperation{op1, op2})
	err := opStorage.Save(op3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	ops, err := opStorage.GetAll()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(ops) != 3 {
		t.Errorf("got ops length: %v, want: %v", len(ops), 3)
	}
	if ops[2] != op3 {
		t.Errorf("got: %v, want: %v", ops[2], op3)
	}
}

func TestOperationStorageStub_DeleteLatest_Test(t *testing.T) {
	op1, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 1000, 1398351430, nil)
	op2, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 500, 1398351431, op1.Hash)
	op3, _ := opCreator.Create(entity.DepositOperation, entity.EUR, 125, 1398351432, op2.Hash)
	opStorage := storage.NewStubOperationStorage([]*entity.CurrencyOperation{op1, op2, op3})
	deleted, err := opStorage.DeleteLatest()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	ops, err := opStorage.GetAll()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(ops) != 2 {
		t.Errorf("got ops length: %v, want: %v", len(ops), 2)
	}
	if deleted != op3 {
		t.Errorf("got: %v, want: %v", ops[2], op3)
	}
}
