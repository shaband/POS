package repo

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
)

type Invoice struct {
	db *bun.DB
}

func NewInvoice(db *bun.DB) *Invoice {
	return &Invoice{db: db}
}

func (r *Invoice) GetInvoices(ctx context.Context) ([]model.Invoice, error) {
	var invoices []model.Invoice

	err := r.db.NewSelect().
		Model(&invoices).
		Relation("Client", selectNameAndID).
		Relation("User", selectNameAndID).
		Relation("Inventory", selectNameAndID).
		Scan(ctx)

	return invoices, err
}

func (r *Invoice) AddInvoice(ctx context.Context, invoiceDTO *dto.InvoiceDTO) (*model.Invoice, error) {
	invoice := model.Invoice{
		InventoryID: invoiceDTO.InventoryID,
		IsSell:      invoiceDTO.IsSell,
		ClientID:    invoiceDTO.ClientID,
		UserID:      invoiceDTO.UserID,
		TotalPrice:  invoiceDTO.TotalPrice,
		TotalCost:   invoiceDTO.TotalPrice,
	}
	q := r.db.NewInsert().
		Model(&invoice)

	_, err := q.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &invoice, err
}

func (r *Invoice) AddItemToInvoice(ctx context.Context, InvoiceID int, invoiceItemDTO *dto.InvoiceItemDTO) (*model.InvoiceItem, error) {
	invoiceItem := model.InvoiceItem{
		InvoiceID:     InvoiceID,
		ProductID:     invoiceItemDTO.ProductID,
		Amount:        invoiceItemDTO.Amount,
		UnitSellPrice: invoiceItemDTO.UnitSellPrice,
		UnitCostPrice: invoiceItemDTO.UnitCostPrice,
		TotalPrice:    invoiceItemDTO.UnitSellPrice * invoiceItemDTO.Amount,
		TotalCost:     invoiceItemDTO.UnitCostPrice * invoiceItemDTO.Amount,
	}
	q := r.db.NewInsert().
		Model(&invoiceItem)
	_, err := q.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &invoiceItem, err
}
