package service

import (
	"context"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/repo"
)

type Client struct {
	Repo *repo.Client
}

func NewClient(Repo *repo.Client) *Client {
	return &Client{
		Repo: Repo,
	}
}

func (s *Client) GetClients(ctx context.Context) ([]model.Client, error) {
	clients, err := s.Repo.GetClients(ctx)
	return clients, err
}

func (s *Client) AddClient(ctx context.Context, clientDTO *dto.CreateClientDTO) (*model.Client, error) {
	return s.Repo.AddClient(ctx, clientDTO)
}

func (s *Client) UpdateClient(ctx context.Context, clientID int, clientDTO *dto.UpdateClientDTO) (*model.Client, error) {
	return s.Repo.UpdateClient(ctx, clientID, clientDTO)
}

func (s *Client) DeleteClient(ctx context.Context, clientID int) error {
	return s.Repo.DeleteClient(ctx, clientID)
}
