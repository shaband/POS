package service

import (
	"context"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/repo"
)

type Category struct {
	repo *repo.Category
}

func NewCategory(repo *repo.Category) *Category {
	return &Category{
		repo: repo,
	}
}

func (s *Category) GetCategories(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *Category) AddCategory(ctx context.Context, categoryDTO *dto.CategoryDTO) (*model.Category, error) {
	return s.repo.AddCategory(ctx, categoryDTO)
}

func (s *Category) UpdateCategory(ctx context.Context, categoryID int, categoryDTO *dto.CategoryDTO) (*model.Category, error) {
	return s.repo.UpdateCategory(ctx, categoryID, categoryDTO)
}

func (s *Category) DeleteCategory(ctx context.Context, categoryID int) error {
	return s.repo.DeleteCategory(ctx, categoryID)
}
