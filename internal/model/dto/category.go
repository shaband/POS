package dto

type CategoryDTO struct {
	Name       string `example:"category1" json:"name"        valid:"type(string),required~name field is required"`
	CategoryId int    `example:"1"         json:"category_id" swaggerignore:"true"                                 valid:"type(int),optional"`
}
