package interactor

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/google/uuid"
	"github.com/hizzuu/plate-backend/internal/domain"
)

func resizeImage(i *domain.Image) error {
	i.Name = uuid.NewString()

	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(i.File); err != nil {
		return err
	}

	img, t, err := image.Decode(buf)
	if err != nil {
		return err
	}

	switch t {
	case "jpeg":
		if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			return err
		}
		i.Name = i.Name + ".jpg"
	case "png":
		if err := png.Encode(buf, img); err != nil {
			return err
		}
		i.Name = i.Name + ".png"
	case "gif":
		if err := gif.Encode(buf, img, nil); err != nil {
			return err
		}
		i.Name = i.Name + ".gif"
	}

	i.File = buf

	return nil
}
