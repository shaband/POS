package dto

type CreateClientDTO struct {
	Name       string `example:"shaband"             json:"name"        valid:"required~name is required,type(string)"`
	Email      string `example:"shaband@shaband.com" json:"email"       valid:"required~name is required,email~email must be vaild email"`
	Phone      string `example:"0123456789"          json:"phone"       valid:"required~phone is required"`
	IsClient   bool   `example:"true"                json:"is_client"   valid:"optional"`
	IsSupplier bool   `example:"true"                json:"is_supplier" valid:"optional"`
	Image      string `json:"image"                  valid:"optional"`
}

type UpdateClientDTO struct {
	Name       string `example:"shaband"             json:"name"        valid:"optional"`
	Email      string `example:"shaband@shaband.com" json:"email"       valid:"optional"`
	Phone      string `example:"0123456789"          json:"phone"       valid:"optional"`
	IsClient   bool   `example:"true"                json:"is_client"   valid:"optional"`
	IsSupplier bool   `example:"true"                json:"is_supplier" valid:"optional"`
	Image      string `json:"image"                  valid:"optional"`
}
