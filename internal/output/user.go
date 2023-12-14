package output

type User struct {
	ID         uint   `example:"1"`
	Name       string `example:"shaband"`
	Email      string `example:"shaband@shaband.com"`
	Phone      string `example:"123456789"`
	isSupplier string `example:"true"`
	isClient   string `example:"false"`
	// Image string `example:"http://pos.test/image.jpg"`
}
