package magick

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	//Files from:
	//- https://github.com/golang/go/issues/2362
	//- https://github.com/golang/go/issues/10447
	//- https://github.com/golang/go/issues/45902
	issues := []string{"2362", "10447", "45902"}
	for _, num := range issues {
		fileName := `./example_images/` + num + `.jpg`
		f1, err := os.Open(fileName)
		if err != nil {
			log.Panic(err)
		}
		defer f1.Close()
		//default way(fail)
		_, _, err = image.Decode(f1)
		assert.Error(t, err)
		log.Printf("Default way error(#%s): %s\n", num, err)

		//Need to open again because reader is changed by `image.Decode()`
		//This library
		f2, err := os.Open(fileName)
		if err != nil {
			log.Panic(err)
		}
		defer f2.Close()
		img, err := Decode(f2)
		assert.NoError(t, err)
		log.Printf("This library succeeded(#%s): e.g. image bounds = %s\n", num, img.Bounds())
	}
}

func ProcessFile() {
	inputFile := ""
	outputFile := ""
	switch len(os.Args) {
	case 1:
		log.Panic("Input image file you want to downsize")
	case 2:
		inputFile = os.Args[1]

		dir := filepath.Dir(inputFile)
		base := filepath.Base(inputFile)
		outputFile = filepath.Join(dir, "_"+base)
	case 3:
		inputFile = os.Args[1]
		outputFile = os.Args[2]
	}

	input, err := os.Open(inputFile)
	if err != nil {
		log.Panic(err)
	}
	defer input.Close()
	img, err := Decode(input)
	if err != nil {
		log.Panic(err)
	}

	imgWriter, err := os.Create(outputFile)
	if err != nil {
		log.Panic(err)
	}
	defer imgWriter.Close()
	jpeg.Encode(imgWriter, img, nil)
}
