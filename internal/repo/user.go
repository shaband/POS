package repo

import (
	"context"
	"errors"

	"github.com/uptrace/bun"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
)

type User struct {
	db *bun.DB
}

func NewUser(db *bun.DB) *User {
	return &User{db: db}
}

func (r *User) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User

	err := r.db.NewSelect().
		Model(&users).
		Scan(ctx)

	return users, err
}

func (r *User) AddUser(ctx context.Context, userDTO *dto.CreateUserDTO) (*model.User, error) {
	user := model.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Phone:    userDTO.Phone,
		Password: userDTO.Password,
	}
	q := r.db.NewInsert().
		Model(&user)
	_, err := q.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *User) UpdateUser(ctx context.Context, userID int, userDTO *dto.UpdateUserDTO) (*model.User, error) {
	user := model.User{
		ID:       userID,
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Phone:    userDTO.Phone,
		Password: userDTO.Password,
	}

	exists, _ := r.db.NewSelect().
		Model(&user).
		WherePK().
		Exists(ctx)
	if !exists {
		return nil, errors.New("User doesn't exists")
	}

	_, err := r.db.NewUpdate().
		Model(&user).
		OmitZero().
		WherePK().
		Exec(ctx)

	if err == nil {
		return &user, err
	}
	return nil, err
}

func (r *User) DeleteUser(ctx context.Context, userID int) error {
	user := model.User{
		ID: userID,
	}

	exists, _ := r.db.NewSelect().
		Model(&user).
		WherePK().
		Exists(ctx)
	if !exists {
		return errors.New("User doesn't exists")
	}

	_, err := r.db.NewDelete().
		Model(&user).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
