package dto

type ProductDTO struct {
	Name       string `form:"name"        json:"name"        valid:"required~name is required"`
	Code       string `form:"code"        json:"code"        valid:"required~code is required"`
	CostPrice  string `form:"cost_price"  json:"cost_price"  valid:"required~cost price is required"`
	SellPrice  string `form:"sell_price"  json:"sell_price"  valid:"required~sell price is required"`
	Image      string `form:"image"       json:"image"       valid:"optional"`
	CategoryID int    `form:"category_id" json:"category_id" swaggerignore:"true"                    valid:"optional"`
}
