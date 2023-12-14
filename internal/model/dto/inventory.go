package dto

type InventoryDTO struct {
	Name    string `example:"inventory1" json:"name"    valid:"type(string),required~name field is required"`
	Address string `example:"address1"   json:"address" valid:"type(string),required~address field is required"`
}
