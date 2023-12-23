package output

type Invoice struct {
	ID         uint `example:"1"`
	Client     NameWithID
	User       NameWithID
	Inventory  NameWithID
	Items      []InvoiceItem
	IsSell     bool
	TotalCost  float64
	TotalPrice float64
}

type InvoiceItem struct {
	ID            uint    `example:"1"`
	UnitCostPrice float64 `example:"2"`
	UnitSellPrice float64 `example:"1"`
	Amount        int     `example:"3"`
	TotalCost     float64 `example:"6"`
	TotalPrice    float64 `example:"3"`
	Product       Product
}

