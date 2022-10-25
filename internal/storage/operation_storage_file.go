package storage

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
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

func (stg *FileOperationStorage) GetAll() ([]*entity.CurrencyOperation, error) {
	f, err := os.Open(stg.filename)
	if err != nil {
		return nil, fmt.Errorf("get list of operations: %v", err)
	}
	defer f.Close()
	ops := make([]*entity.CurrencyOperation, 0)
	rdr := bufio.NewReader(f)
	for {
		line, err := rdr.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("get list of operations: %v", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			return ops, nil
		}
		op, err := stg.OperationFromString(line)
		if err != nil {
			return nil, fmt.Errorf("get list of operations: %v", err)
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (stg *FileOperationStorage) GetLatest() (*entity.CurrencyOperation, error) {
	makeError := func(err error) error {
		return fmt.Errorf("get latest operation: %v", err)
	}
	f, err := os.Open(stg.filename)
	if err != nil {
		return nil, makeError(err)
	}
	defer f.Close()
	off, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}
	size := int64(256)
	for {
		buf := make([]byte, size)
		newOff := off - size
		if newOff < 0 {
			newOff = 0
		}
		n, _ := f.ReadAt(buf, newOff)
		lines := bytes.Fields(buf[:n])
		if newOff == 0 && len(lines) == 0 {
			return nil, entity.ErrNoOperations
		}
		if len(lines) > 0 {
			last := lines[len(lines)-1]
			op, err := stg.OperationFromString(string(last))
			if err != nil {
				return nil, makeError(err)
			}
			return op, nil
		}
		size *= 2
	}
}

func (stg *FileOperationStorage) Save(op *entity.CurrencyOperation) error {
	f, err := os.OpenFile(stg.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("save operation: %v", err)
	}
	defer f.Close()
	_, err = f.WriteString(stg.OperationToString(op))
	if err != nil {
		return fmt.Errorf("save operation: %v", err)
	}
	return nil
}

func (stg *FileOperationStorage) DeleteLatest() (*entity.CurrencyOperation, error) {
	makeError := func(err error) error {
		return fmt.Errorf("delete latest operation: %v", err)
	}
	f, err := os.Open(stg.filename)
	if err != nil {
		return nil, makeError(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, makeError(err)
	}
	data = data[:len(data)-1] // cut last line break (\n)
	idx := bytes.LastIndex(data, []byte{'\n'})
	if err := os.WriteFile(stg.filename, data[:idx+1], 0666); err != nil {
		return nil, makeError(err)
	}
	op, err := stg.OperationFromString(string(data[idx+1:]))
	if err != nil {
		return nil, makeError(err)
	}
	return op, nil
}

func (stg *FileOperationStorage) OperationFromString(s string) (*entity.CurrencyOperation, error) {
	fmtError := func(err error) error {
		return fmt.Errorf("parse operation string: %v", err)
	}
	parts := strings.Split(strings.TrimSuffix(s, "\n"), ",")
	if len(parts) != 6 {
		return nil, fmtError(fmt.Errorf("want 6 parts, got: %v, %#v", len(parts), parts))
	}
	opType, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, fmtError(err)
	}
	currency, err := entity.CurrencyFromString(parts[1])
	if err != nil {
		return nil, fmtError(err)
	}
	amount, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return nil, fmtError(err)
	}
	providedAt, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return nil, fmtError(err)
	}
	hash, err := hex.DecodeString(parts[4])
	if err != nil {
		return nil, fmtError(err)
	}
	prevHash, err := hex.DecodeString(parts[5])
	if err != nil {
		return nil, fmtError(err)
	}
	op := entity.NewCurrencyOperation(entity.OperationType(opType), currency, amount, providedAt, hash, prevHash)
	return op, nil
}

func (stg *FileOperationStorage) OperationToString(op *entity.CurrencyOperation) string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", op.Optype, op.Currency, op.Amount, op.ProvidedAt, hex.EncodeToString(op.Hash), hex.EncodeToString(op.PrevHash))
}
