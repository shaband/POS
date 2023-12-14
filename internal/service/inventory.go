package service

import (
	"context"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/repo"
)

type Inventory struct {
	repo *repo.Inventory
}

func NewInventory(repo *repo.Inventory) *Inventory {
	return &Inventory{
		repo: repo,
	}
}

func (s *Inventory) GetInventories(ctx context.Context) ([]model.Inventory, error) {
	return s.repo.GetInventories(ctx)
}

func (s *Inventory) AddInventory(ctx context.Context, inventoryDTO *dto.InventoryDTO) (*model.Inventory, error) {
	return s.repo.AddInventory(ctx, inventoryDTO)
}

func (s *Inventory) UpdateInventory(ctx context.Context, inventoryID int, inventoryDTO *dto.InventoryDTO) (*model.Inventory, error) {
	return s.repo.UpdateInventory(ctx, inventoryID, inventoryDTO)
}

func (s *Inventory) DeleteInventory(ctx context.Context, inventoryID int) error {
	return s.repo.DeleteInventory(ctx, inventoryID)
}
