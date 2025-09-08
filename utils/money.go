package utils

import (
	"errors"
	"fmt"
)

type Money struct {
	amount int64
}

func NewMoney(cents int64) (Money, error) {
	if cents < 0 {
		return Money{}, errors.New("money amount cannot be negative")
	}
	return Money{amount: cents}, nil
}

func NewMoneyFromDollars(dollars float64) (Money, error) {
	if dollars < 0 {
		return Money{}, errors.New("money amount cannot be negative")
	}
	cents := int64(dollars * 100)
	return Money{amount: cents}, nil
}

func (m Money) Cents() int64 {
	return m.amount
}

func (m Money) Dollars() float64 {
	return float64(m.amount) / 100
}

func (m Money) String() string {
	return fmt.Sprintf("$%.2f", m.Dollars())
}

func (m Money) Add(other Money) Money {
	return Money{amount: m.amount + other.amount}
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.amount < other.amount {
		return Money{}, errors.New("insufficient funds")
	}
	return Money{amount: m.amount - other.amount}, nil
}

func (m Money) GreaterThanOrEqual(other Money) bool {
	return m.amount >= other.amount
}

func (m Money) IsZero() bool {
	return m.amount == 0
}

func (m Money) IsPositive() bool {
	return m.amount > 0
}
