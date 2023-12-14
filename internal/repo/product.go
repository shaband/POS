package repo

import (
	"context"
	"errors"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/uptrace/bun"
)

type Product struct {
	db *bun.DB
}

func NewProduct(db *bun.DB) *Product {
	return &Product{
		db: db,
	}
}

func (r Product) GetProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	err := r.db.NewSelect().
		Model(&products).
		Scan(ctx)

	return products, err
}

func (r Product) AddProduct(ctx context.Context, productDTO *dto.ProductDTO) (*model.Product, error) {
	product := &model.Product{
		Name:       productDTO.Name,
		Code:       productDTO.Code,
		CostPrice:  productDTO.CostPrice,
		SellPrice:  productDTO.SellPrice,
		CategoryID: productDTO.CategoryID,
		Image:      productDTO.Image,
	}
	_, err := r.db.NewInsert().Model(product).Exec(ctx)

	return product, err
}

func (r Product) UpdateProduct(ctx context.Context, productID int, productDTO *dto.ProductDTO) (*model.Product, error) {
	product := &model.Product{
		ID:         productID,
		Name:       productDTO.Name,
		Code:       productDTO.Code,
		CostPrice:  productDTO.CostPrice,
		SellPrice:  productDTO.SellPrice,
		CategoryID: productDTO.CategoryID,
		Image:      productDTO.Image,
	}

	exists, _ := r.db.NewSelect().
		Model(product).
		WherePK().
		Exists(ctx)
	if !exists {
		return nil, errors.New("Product doesn't exists")
	}
	_, err := r.db.NewUpdate().Model(&product).WherePK().OmitZero().Exec(ctx)

	return product, err
}

func (r *Product) DeleteProduct(ctx context.Context, productID int) error {
	product := &model.Product{
		ID: productID,
	}

	exists, _ := r.db.NewSelect().
		Model(product).
		WherePK().
		Exists(ctx)
	if !exists {
		return errors.New("Product doesn't exists")
	}

	_, err := r.db.NewDelete().
		Model(product).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
