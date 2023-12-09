package repo

import (
	"context"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/uptrace/bun"
)

type Category struct {
	db *bun.DB
}

func NewCategory(db *bun.DB) *Category {
	return &Category{
		db: db,
	}
}

func (r *Category) GetCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category

	err := r.db.NewSelect().
		Model(&categories).
		Scan(ctx)

	return categories, err
}

func (r *Category) AddCategory(ctx context.Context, categoryDTO *dto.CategoryDTO) (*model.Category, error) {
	category := model.Category{
		Name:       categoryDTO.Name,
		CategoryID: categoryDTO.CategoryId,
	}
	q := r.db.NewInsert().
		Model(&category)
	if category.CategoryID == 0 {
		q.ExcludeColumn("category_id")
	}
	result, err := q.Exec(ctx)
	if err != nil {
		return nil, err
	}
	LastID, _ := result.LastInsertId()
	category.ID = int(LastID)
	return &category, err
}

func (r *Category) UpdateCategory(ctx context.Context, categoryID int, categoryDTO *dto.CategoryDTO) (*model.Category, error) {
	category := model.Category{
		ID:         categoryID,
		Name:       categoryDTO.Name,
		CategoryID: categoryDTO.CategoryId,
	}
	q := r.db.NewUpdate().
		Model(&category).
		WherePK()
	if category.CategoryID == 0 {
		q.ExcludeColumn("category_id")
	}
	_, err := q.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &category, err
}

func (r *Category) DeleteCategory(ctx context.Context, categoryID int) error {
	category := model.Category{
		ID: categoryID,
	}

	_, err := r.db.NewDelete().
		Model(&category).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
