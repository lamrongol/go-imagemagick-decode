package magick

import (
	"bytes"
	"image"
	"io"
	"os/exec"
)

// Alternative of `image.Decode()`
// You need to install ImageMagick( https://imagemagick.org/ )
// based on this Japanese article https://zenn.dev/aeon_mall/articles/b54d3abad707d8
func Decode(img io.Reader) (image.Image, error) {
	//Hold original data because `image.Decode()` changes reader point
	original, err := io.ReadAll(img)
	if err != nil {
		return nil, err
	}

	//First, decode by default way because it takes time to run ImageMagick
	decodeImage, _, err := image.Decode(img)
	//If default way is failed
	if err != nil {
		downsizedImage, err := downsizeSamplingFactor(original)
		if err != nil {
			return nil, err
		}
		decodeImage, _, err = image.Decode(bytes.NewReader(downsizedImage))
		if err != nil {
			return nil, err
		}
	}
	return decodeImage, nil
}

func downsizeSamplingFactor(img []byte) ([]byte, error) {
	params := []string{
		"-",                         //Stdin image
		"-sampling-factor", "4:4:4", //downsize Chroma Sub-sampling
		"-", //Stdout image
	}
	cmd := exec.Command("magick", params...)
	cmd.Stdin = bytes.NewReader(img)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
