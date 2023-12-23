package service

import (
	"context"
	"reflect"

	"github.com/samber/lo"

	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/repo"
)

type Invoice struct {
	Repo         *repo.Invoice
	ProductPrice *repo.Product
}

func NewInvoice(Repo *repo.Invoice) *Invoice {
	return &Invoice{
		Repo: Repo,
	}
}

func (s *Invoice) GetInvoices(ctx context.Context) ([]model.Invoice, error) {
	invoices, err := s.Repo.GetInvoices(ctx)
	return invoices, err
}

func (s *Invoice) AddInvoice(ctx context.Context, invoiceDTO *dto.InvoiceDTO) (*model.Invoice, error) {

if(invoiceDTO.IsSell){
	
}
	invoiceDTO.TotalCost = SumTotalForInvoice(invoiceDTO.Items, "TotalCost")
	invoiceDTO.TotalPrice = SumTotalForInvoice(invoiceDTO.Items, "TotalPrice")

	return s.Repo.AddInvoice(ctx, invoiceDTO)
}

func SumTotalForInvoice(Items []dto.InvoiceItemDTO, field string) float64 {
	return lo.Reduce(Items, func(agg float64, item dto.InvoiceItemDTO, index int) float64 {
		r := reflect.ValueOf(item)
		f := reflect.Indirect(r).FieldByName(field)
		return agg + f.Float()
	}, 0)
}
