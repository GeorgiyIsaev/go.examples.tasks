package u04_full

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(order *Order) error {
	args := m.Called(order)
	return args.Error(0)
}

type MockPayment struct {
	mock.Mock
}

func (m *MockPayment) Charge(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}

func TestService_CreateOrder_Success(t *testing.T) {
	repo := new(MockRepo)
	payment := new(MockPayment)

	// Ожидания
	payment.On("Charge", 100.0).Return(nil)
	repo.On("Save", mock.AnythingOfType("*u04_full.Order")).Return(nil)

	svc := NewService(repo, payment)
	err := svc.CreateOrder(100.0)

	assert.NoError(t, err)
	payment.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestService_CreateOrder_PaymentFails(t *testing.T) {
	repo := new(MockRepo)
	payment := new(MockPayment)

	payment.On("Charge", 100.0).Return(errors.New("payment declined"))
	// Не ожидаем вызова Save

	svc := NewService(repo, payment)
	err := svc.CreateOrder(100.0)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "payment declined")
	payment.AssertExpectations(t)
	repo.AssertExpectations(t) // Save не вызывался — тест пройден
}
