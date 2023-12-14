package helper

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Uploader struct{}

func RegisterUploader() *Uploader {
	return &Uploader{}
}

func (e *Uploader) SaveFIle(c *fiber.Ctx, inputName, prefix string) (string, error) {
	file, err := c.FormFile(inputName)
	log.Err(err).Msg("error on get file")
	if err != nil {
		return "", err
	}
	path := "/storage/uploads/" + prefix + "_" + fmt.Sprint(time.Now().Unix()) + "_" + file.Filename
	// Save file to root directory:
	err = c.SaveFile(file, "."+path)

	if err != nil {
		return "", err
	}
	return path, nil
}
