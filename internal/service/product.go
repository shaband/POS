package service

import (
	"context"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/repo"
)

type Product struct {
	Repo *repo.Product
}

func NewProduct(Repo *repo.Product) *Product {
	return &Product{
		Repo: Repo,
	}
}

func (s *Product) GetProducts(ctx context.Context) ([]model.Product, error) {
	products, err := s.Repo.GetProducts(ctx)
	return products, err
}

func (s *Product) AddProduct(ctx context.Context, productDTO *dto.ProductDTO) (*model.Product, error) {
	return s.Repo.AddProduct(ctx, productDTO)
}

func (s *Product) ShowProduct(ctx context.Context, productID int) (*model.Product, error) {
	return s.Repo.ShowProduct(ctx, productID)
}

func (s *Product) UpdateProduct(ctx context.Context, productID int, productDTO *dto.ProductDTO) (*model.Product, error) {
	return s.Repo.UpdateProduct(ctx, productID, productDTO)
}

func (s *Product) DeleteProduct(ctx context.Context, productID int) error {
	return s.Repo.DeleteProduct(ctx, productID)
}
