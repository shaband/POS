package service

import (
	"context"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/repo"
)

type User struct {
	Repo *repo.User
}

func NewUser(Repo *repo.User) *User {
	return &User{
		Repo: Repo,
	}
}

func (s *User) GetUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.Repo.GetUsers(ctx)
	return users, err
}

func (s *User) AddUser(ctx context.Context, userDTO *dto.CreateUserDTO) (*model.User, error) {
	return s.Repo.AddUser(ctx, userDTO)
}

func (s *User) UpdateUser(ctx context.Context, userID int, userDTO *dto.UpdateUserDTO) (*model.User, error) {
	return s.Repo.UpdateUser(ctx, userID, userDTO)
}

func (s *User) DeleteUser(ctx context.Context, userID int) error {
	return s.Repo.DeleteUser(ctx, userID)
}
