package service

import (
	"context"
	"reflect"
	"strconv"

	"github.com/samber/lo"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/repo"
)

type Invoice struct {
	Repo          *repo.Invoice
	InventoryRepo *repo.Inventory
	ProductRepo   *repo.Product
}

func NewInvoice(Repo *repo.Invoice, InventoryRepo *repo.Inventory, ProductRepo *repo.Product) *Invoice {
	return &Invoice{
		Repo:          Repo,
		InventoryRepo: InventoryRepo,
		ProductRepo:   ProductRepo,
	}
}

func (s *Invoice) GetInvoices(ctx context.Context) ([]model.Invoice, error) {
	invoices, err := s.Repo.GetInvoices(ctx)
	return invoices, err
}

func (s *Invoice) AddInvoice(ctx context.Context, invoiceDTO *dto.InvoiceDTO) (*model.Invoice, error) {
	items, _ := s.getCostAndPricePerItem(ctx, invoiceDTO.IsSell, invoiceDTO.Items)

	invoiceDTO.TotalCost = SumTotalForInvoice(items, "UnitCostPrice")
	invoiceDTO.TotalPrice = SumTotalForInvoice(items, "UnitSellPrice")

	invoice, err := s.Repo.AddInvoice(ctx, invoiceDTO)
	if err != nil {
		return nil, err
	}
	s.addInvoiceItems(ctx, invoice.ID, invoice.InventoryID, invoice.IsSell, items)
	return &model.Invoice{}, nil
}

func SumTotalForInvoice(Items []dto.InvoiceItemDTO, field string) float64 {
	return lo.Reduce(Items, func(agg float64, item dto.InvoiceItemDTO, index int) float64 {
		r := reflect.ValueOf(item)

		f := reflect.Indirect(r).FieldByName(field)
		return agg + f.Float()
	}, 0)
}

func (s *Invoice) addInvoiceItems(ctx context.Context, InvoiceID, inventoryID int, isSell bool, items []dto.InvoiceItemDTO) {
	for _, item := range items {
		s.Repo.AddItemToInvoice(ctx, InvoiceID, &item)
		inventoryToProductDTO := dto.InventoryToProductDTO{
			ProductID:   item.ProductID,
			InventoryID: inventoryID,
			Amount:      int(item.Amount),
		}
		if isSell {
			s.InventoryRepo.SubFromInventory(ctx, inventoryID, &inventoryToProductDTO)
		} else {
			s.InventoryRepo.AddToInventory(ctx, inventoryID, &inventoryToProductDTO)
			s.ProductRepo.UpdateProductPrices(ctx, item.ProductID, item.UnitCostPrice, item.UnitSellPrice)
		}
	}
}

func (s *Invoice) getCostAndPricePerItem(ctx context.Context, isSell bool, items []dto.InvoiceItemDTO) ([]dto.InvoiceItemDTO, error) {
	if !isSell {
		return items, nil
	}

	if isSell {
		productIDs := []int{}
		for _, prod := range items {
			productIDs = append(productIDs, prod.ProductID)
		}

		products, err := s.ProductRepo.GetProductsByIDS(ctx, &productIDs)
		if err != nil {
			return nil, err
		}
		newItems := []dto.InvoiceItemDTO{}
		for _, i := range items {
			result, _ := lo.Find(products, func(item model.Product) bool {
				return item.ID == i.ProductID
			})

			cost, _ := strconv.ParseFloat(result.CostPrice, 64)
			i.UnitCostPrice = cost
			sellPrice, _ := strconv.ParseFloat(result.SellPrice, 64)
			i.UnitSellPrice = sellPrice
			newItems = append(newItems, i)
		}
		return newItems, nil
	}
	return items, nil
}
