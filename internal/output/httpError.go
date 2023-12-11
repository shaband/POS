package output

type HTTPError struct {
	Success bool   `example:"false"`
	Message string `example:"name field is required"`
}
