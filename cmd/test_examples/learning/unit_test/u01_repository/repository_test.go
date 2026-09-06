package u01_repository

import "testing"

// Стаб-реализация
type stubUserRepository struct {
	users map[int]*User
}

func (r stubUserRepository) GetByID(id int) (*User, error) {
	if u, ok := r.users[id]; ok {
		return u, nil
	}
	return nil, ErrNotFound
}

func (r stubUserRepository) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

func TestGetUser(t *testing.T) {
	repo := stubUserRepository{
		users: map[int]*User{
			1: {ID: 1, Name: "Alice"},
		},
	}
	service := NewUserService(repo)

	user, err := service.GetUser(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Name != "Alice" {
		t.Errorf("expected Alice, got %s", user.Name)
	}
}
