package u02_gomock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetUserWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)
	expectedUser := &User{ID: 1, Name: "Bob"}
	mockRepo.EXPECT().
		GetByID(1).
		Return(expectedUser, nil)

	service := NewUserService(mockRepo)
	user, err := service.GetUser(1)
	_, _ = user, err
	// проверки...
}
