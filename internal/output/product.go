package output

type Product struct {
	ID        int    `example:"1"                          json:"id"`
	Name      string `example:"product 1"                  json:"name"`
	Code      string `example:"124"                        json:"code"`
	CostPrice string `example:"5"                          json:"cost_price"`
	SellPrice string `example:"10"                         json:"sell_price"`
	Image     string `example:"http://image.com/image.jpg" json:"image"`
}
type ProductInventory struct {
	ID                uint   `example:"1"         json:"id"`
	Name              string `example:"inventory" json:"name"`
	Address           string `example:"address1"  json:"address"`
	Amount            int    `example:"1"         json:"amount"`
	SellInvoicesCount int    `example:"0"         json:"sell_invoices_count"`
	BuyInvoicesCount  int    `example:"0"         json:"buy_invoices_count"`
}
type ProductDetails struct {
	ID          int    `example:"1"                          json:"id"`
	Name        string `example:"product 1"                  json:"name"`
	Code        string `example:"124"                        json:"code"`
	CostPrice   string `example:"5"                          json:"cost_price"`
	SellPrice   string `example:"10"                         json:"sell_price"`
	Image       string `example:"http://image.com/image.jpg" json:"image"`
	Category    Category
	Inventories []ProductInventory
}
