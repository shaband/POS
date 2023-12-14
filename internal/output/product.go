package output

type Product struct {
	ID        int    `example:"1"                          json:"id"`
	Name      string `example:"product 1"                  json:"name"`
	Code      string `example:"124"                        json:"code"`
	CostPrice string `example:"5"                          json:"cost_price"`
	SellPrice string `example:"10"                         json:"sell_price"`
	Image     string `example:"http://image.com/image.jpg" json:"image"`
}
