package repo

import (
	"context"
	"errors"

	"github.com/uptrace/bun"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
)

type Client struct {
	db *bun.DB
}

func NewClient(db *bun.DB) *Client {
	return &Client{db: db}
}

func (r *Client) GetClients(ctx context.Context) ([]model.Client, error) {
	var clients []model.Client

	err := r.db.NewSelect().
		Model(&clients).
		Scan(ctx)

	return clients, err
}

func (r *Client) AddClient(ctx context.Context, clientDTO *dto.CreateClientDTO) (*model.Client, error) {
	client := model.Client{
		Name:       clientDTO.Name,
		Email:      clientDTO.Email,
		Phone:      clientDTO.Phone,
		IsSupplier: clientDTO.IsSupplier,
		IsClient:   clientDTO.IsClient,
	}
	q := r.db.NewInsert().
		Model(&client)
	_, err := q.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &client, err
}

func (r *Client) UpdateClient(ctx context.Context, clientID int, clientDTO *dto.UpdateClientDTO) (*model.Client, error) {
	client := model.Client{
		ID:         clientID,
		Name:       clientDTO.Name,
		Email:      clientDTO.Email,
		Phone:      clientDTO.Phone,
		IsSupplier: clientDTO.IsSupplier,
		IsClient:   clientDTO.IsClient,
	}

	exists, _ := r.db.NewSelect().
		Model(&client).
		WherePK().
		Exists(ctx)
	if !exists {
		return nil, errors.New("Client doesn't exists")
	}

	_, err := r.db.NewUpdate().
		Model(&client).
		OmitZero().
		WherePK().
		Exec(ctx)

	if err == nil {
		return &client, err
	}
	return nil, err
}

func (r *Client) DeleteClient(ctx context.Context, clientID int) error {
	client := model.Client{
		ID: clientID,
	}

	exists, _ := r.db.NewSelect().
		Model(&client).
		WherePK().
		Exists(ctx)
	if !exists {
		return errors.New("Client doesn't exists")
	}

	_, err := r.db.NewDelete().
		Model(&client).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
