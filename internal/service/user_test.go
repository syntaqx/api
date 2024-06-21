package service

import (
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"

	"github.com/syntaqx/api/internal/model"
	"github.com/syntaqx/api/internal/repository/mock"
)

func TestUserServiceCRUD(t *testing.T) {
	users := []*model.User{}
	repository := &mock.UserRepositoryMock{
		CreateUserFunc: func(user *model.User) error {
			users = append(users, user)
			return nil
		},
		ListUsersFunc: func() ([]*model.User, error) {
			return users, nil
		},
		GetUserByIDFunc: func(id uuid.UUID) (*model.User, error) {
			for _, user := range users {
				if user.ID == id {
					return user, nil
				}
			}
			return nil, nil
		},
		DeleteUserFunc: func(id uuid.UUID) error {
			for i, user := range users {
				if user.ID == id {
					users = append(users[:i], users[i+1:]...)
					return nil
				}
			}
			return nil
		},
	}

	// Create a new UserService instance
	userService := NewUserService(repository)

	uuid := uuid.Must(uuid.NewV4())

	user := &model.User{
		ID:    uuid,
		Login: "test",
		Email: "test@test.locahost",
	}

	err := userService.CreateUser(user)

	assert.NoError(t, err, "CreateUser should not return an error")
	assert.Equal(t, uuid, user.ID, "User ID should be the same as the one provided")

	users, err = userService.ListUsers()
	assert.NoError(t, err, "ListUsers should not return an error")
	assert.Len(t, users, 1, "ListUsers should return 1 user")

	actualUser, err := userService.GetUserByID(user.ID)
	assert.NoError(t, err, "GetUserByID should not return an error")
	assert.Equal(t, user, actualUser, "GetUserByID should return the correct user")

	err = userService.DeleteUser(user.ID)
	assert.NoError(t, err, "DeleteUser should not return an error")

	users, err = userService.ListUsers()
	assert.NoError(t, err, "ListUsers should not return an error")
	assert.Len(t, users, 0, "ListUsers should return 0 users")
}

func TestCreateUserWithoutUUID(t *testing.T) {
	repository := &mock.UserRepositoryMock{
		CreateUserFunc: func(user *model.User) error {
			return nil
		},
	}
	userService := NewUserService(repository)

	user := &model.User{
		Login: "test",
	}

	err := userService.CreateUser(user)
	assert.NoError(t, err, "CreateUser should not return an error")
	assert.NotEmpty(t, user.ID, "User ID should not be empty")
}
