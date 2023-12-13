package output

type Client struct {
	Id         uint   `example:"1"`
	Name       string `example:"shaband"`
	Email      string `example:"shaband@shaband.com"`
	Phone      string `example:"123456789"`
	IsClient   bool   `example:"true"`
	IsSupplier bool   `example:"true"`
	// Image string `example:"http://pos.test/image.jpg"`
}
