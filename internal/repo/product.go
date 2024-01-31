package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
)

type Product struct {
	db *bun.DB
}

func NewProduct(db *bun.DB) *Product {
	return &Product{
		db: db,
	}
}

func (r *Product) GetProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	err := r.db.NewSelect().
		Model(&products).
		Scan(ctx)

	return products, err
}

func (r *Product) AddProduct(ctx context.Context, productDTO *dto.ProductDTO) (*model.Product, error) {
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

func (r *Product) ShowProduct(ctx context.Context, productID int) (*model.Product, error) {
	product := &model.Product{
		ID: productID,
	}
	err := r.db.NewSelect().
		Model(product).
		WherePK().
		Relation("Category").
		Relation("Inventories",
			func(q1 *bun.SelectQuery) *bun.SelectQuery {
				return q1.Relation("ProductStock")
			}).Scan(ctx)
	// err = nil
	return product, err
}

func (r *Product) UpdateProduct(ctx context.Context, productID int, productDTO *dto.ProductDTO) (*model.Product, error) {
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
	_, err := r.db.NewUpdate().Model(product).WherePK().OmitZero().Exec(ctx)

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

func (r *Product) GetProductsByIDS(ctx context.Context, productIDS *[]int) ([]model.Product, error) {
	products := []model.Product{}
	for _, id := range *productIDS {
		products = append(products, model.Product{
			ID: id,
		})
	}

	err := r.db.NewSelect().Model(&products).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return products, err
}

func (r Product) UpdateProductPrices(ctx context.Context, productID int, cost, price float64) error {
	product := &model.Product{
		ID:        productID,
		CostPrice: fmt.Sprintf("%g", cost),
		SellPrice: fmt.Sprintf("%g", price),
	}
	_, err := r.db.NewUpdate().Model(product).WherePK().OmitZero().Exec(ctx)

	return err
}
