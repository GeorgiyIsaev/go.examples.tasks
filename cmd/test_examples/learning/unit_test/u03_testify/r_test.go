package u03_testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByID(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func TestGetUserWithTestify(t *testing.T) {
	mockRepo := new(MockUserRepository)
	expectedUser := &User{ID: 1, Name: "Bob"}
	mockRepo.On("GetByID", 1).Return(expectedUser, nil)

	service := NewUserService(mockRepo)
	user, err := service.GetUser(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t) // проверяет, что все ожидания выполнены
}
