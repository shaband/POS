package dto

type ProductDTO struct {
	Name       string `json:"name"        form:"name"                             valid:"required~name is required"`
	Code       string `json:"code"        form:"code"                             valid:"required~code is required"`
	CostPrice  string `json:"cost_price"  form:"cost_price"                       valid:"required~cost price is required"`
	SellPrice  string `json:"sell_price"  form:"sell_price"                       valid:"required~sell price is required"`
	Image      string `json:"image"       form:"image" valid:"optional"`
	CategoryID int    `json:"category_id" form:"category_id" swaggerignore:"true" valid:"optional"`
}
