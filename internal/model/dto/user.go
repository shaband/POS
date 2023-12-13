package dto

type CreateUserDTO struct {
	Name                 string `example:"shaband"             json:"name"                  valid:"required,type(string)"`
	Email                string `example:"shaband@shaband.com" json:"email"                 valid:"required,email"`
	Phone                string `example:"0123456789"          json:"phone"                 valid:"required"`
	Password             string `example:"123"                 json:"password"              valid:"required"`
	PasswordConfirmation string `example:"123"                 json:"password_confirmation" valid:"required"`
	// Image                multipart.File `json:"image"                  valid:"optional"`
}

type UpdateUserDTO struct {
	Name                 string `example:"shaband"             json:"name"                  valid:"optional"`
	Email                string `example:"shaband@shaband.com" json:"email"                 valid:"optional"`
	Phone                string `example:"0123456789"          json:"phone"                 valid:"optional"`
	Password             string `example:"123"                 json:"password"              valid:"optional"`
	PasswordConfirmation string `example:"123"                 json:"password_confirmation" valid:"optional"`
	// Image                multipart.File `json:"image"                  valid:"optional"`
}

