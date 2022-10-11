package storage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type FileOperationStorage struct {
	filename string
}

func NewFileOperationStorage(filename string) *FileOperationStorage {
	exe, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("loading operations from file %v: %v", filename, err))
	}
	return &FileOperationStorage{
		filename: path.Join(path.Dir(exe), filename),
	}
}

func (stg *FileOperationStorage) List() ([]*domain.CurrencyOperation, error) {
	f, err := os.Open(stg.filename)
	if err != nil {
		return nil, fmt.Errorf("get list of operations: %v", err)
	}
	defer f.Close()
	var ops []*domain.CurrencyOperation
	rdr := bufio.NewReader(f)
	for {
		line, err := rdr.ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("get list of operations: %v", err)
		}
		op, err := stg.operationFromString(strings.TrimSpace(line))
		if err != nil {
			return nil, fmt.Errorf("get list of operations: %v", err)
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (stg *FileOperationStorage) Save(op *domain.CurrencyOperation) error {
	f, err := os.OpenFile(stg.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("save operation: %v", err)
	}
	_, err = f.WriteString(stg.operationToString(op))
	if err != nil {
		return fmt.Errorf("save operation: %v", err)
	}
	// if err := os.WriteFile(stg.filename, []byte(stg.operationToString(op)), 0644); err != nil {
	// 	return fmt.Errorf("save operation: %v", err)
	// }
	return nil
}

func (stg *FileOperationStorage) DeleteLast() (*domain.CurrencyOperation, error) {
	return nil, nil
}

func (stg *FileOperationStorage) operationFromString(s string) (*domain.CurrencyOperation, error) {
	parts := strings.Split(strings.TrimSuffix(s, ";"), ",")
	if len(parts) != 4 {
		return nil, fmt.Errorf("parse operation string, bad parts: %#v", parts)
	}
	opType, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse operation string: %v", err)
	}
	currency := domain.Currency(parts[1])
	amount, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return nil, fmt.Errorf("parse operation string: %v", err)
	}
	providedAt, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse operation string: %v", err)
	}
	op := domain.NewCurrencyOperation(domain.OperationType(opType), currency, amount, providedAt)
	return op, nil
}

func (stg *FileOperationStorage) operationToString(op *domain.CurrencyOperation) string {
	return fmt.Sprintf("%v,%v,%v,%v;", op.Optype, op.Currency, op.Amount, op.ProvidedAt)
}
