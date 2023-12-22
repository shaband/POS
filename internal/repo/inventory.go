package repo

import (
	"context"
	"errors"

	"github.com/uptrace/bun"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
)

type Inventory struct {
	db *bun.DB
}

func NewInventory(db *bun.DB) *Inventory {
	return &Inventory{
		db: db,
	}
}

func (r *Inventory) GetInventories(ctx context.Context) ([]model.Inventory, error) {
	var inventories []model.Inventory

	err := r.db.NewSelect().
		Model(&inventories).
		Scan(ctx)

	return inventories, err
}

func (r *Inventory) AddInventory(ctx context.Context, inventoryDTO *dto.InventoryDTO) (*model.Inventory, error) {
	inventory := model.Inventory{
		Name:    inventoryDTO.Name,
		Address: inventoryDTO.Address,
	}
	_, err := r.db.NewInsert().
		Model(&inventory).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &inventory, err
}

func (r *Inventory) UpdateInventory(ctx context.Context, inventoryID int, inventoryDTO *dto.InventoryDTO) (*model.Inventory, error) {
	inventory := model.Inventory{
		ID:      inventoryID,
		Name:    inventoryDTO.Name,
		Address: inventoryDTO.Address,
	}
	exists, _ := r.db.NewSelect().
		Model(&inventory).
		WherePK().
		Exists(ctx)
	if !exists {
		return nil, errors.New("Inventory doesn't exists")
	}

	_, err := r.db.NewUpdate().
		Model(&inventory).
		OmitZero().
		WherePK().
		Exec(ctx)

	if err == nil {
		return &inventory, err
	}
	return nil, err
}

func (r *Inventory) DeleteInventory(ctx context.Context, inventoryID int) error {
	inventory := model.Inventory{
		ID: inventoryID,
	}

	exists, _ := r.db.NewSelect().
		Model(&inventory).
		WherePK().
		Exists(ctx)
	if !exists {
		return errors.New("Inventory doesn't exists")
	}

	_, err := r.db.NewDelete().
		Model(&inventory).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
