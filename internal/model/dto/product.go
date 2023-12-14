package dto

type ProductDTO struct {
	Name       string `json:"name"        valid:"required~name is required"`
	Code       string `json:"code"        valid:"required~code is required"`
	CostPrice  string `json:"cost_price"  valid:"required~cost price is required"`
	SellPrice  string `json:"sell_price"  valid:"required"`
	Image      string `json:"image"       valid:"optional"`
	CategoryID int    `json:"category_id" swaggerignore:"true"                    vaild:"optional"`
}
