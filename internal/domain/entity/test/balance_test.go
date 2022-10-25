package test

import (
	"testing"

	"github.com/dmitruk-v/piggy-bank/internal/domain/entity"
)

func TestBalanceAdd(t *testing.T) {
	balance := entity.NewBalanceImpl([]entity.Currency{entity.EUR, entity.UAH})
	if err := balance.Add(entity.EUR, 100); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	got := balance.Get(entity.EUR)
	if got != 100 {
		t.Errorf("got: %v, want: %v", got, 100)
	}
	if err := balance.Add(entity.EUR, 50); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	got = balance.Get(entity.EUR)
	if got != 150 {
		t.Errorf("got: %v, want: %v", got, 150)
	}
}

func TestBalanceAddError(t *testing.T) {
	balance := entity.NewBalanceImpl([]entity.Currency{entity.EUR, entity.UAH})
	if err := balance.Add(entity.Currency(-1), 100); err == nil {
		t.Errorf("got %v, expected error", err)
	}
}

func TestBalanceSub(t *testing.T) {
	balance := entity.NewBalanceImpl([]entity.Currency{entity.USD, entity.UAH})
	if err := balance.Add(entity.USD, 100); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := balance.Sub(entity.USD, 50); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	val := balance.Get(entity.USD)
	if val != 50 {
		t.Errorf("got %v, want %v", val, 50)
	}
	if err := balance.Sub(entity.USD, 50); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	val = balance.Get(entity.USD)
	if val != 0 {
		t.Errorf("got %v, want %v", val, 0)
	}
}

func TestBalanceSubError(t *testing.T) {
	balance := entity.NewBalanceImpl([]entity.Currency{entity.USD, entity.UAH})
	if err := balance.Add(entity.USD, 100); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := balance.Sub(entity.Currency(-1), 50); err == nil {
		t.Errorf("got %v, expected error", err)
	}
	if err := balance.Sub(entity.USD, 101); err == nil {
		t.Errorf("got %v, expect error", err)
	}
}
