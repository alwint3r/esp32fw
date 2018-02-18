package esp32fw

import (
	"bufio"
	"bytes"
	"errors"
	"os"
)

// ErrEmptyOutputPath Error returned when output path to a firmware is empty
var ErrEmptyOutputPath = errors.New("output path is empty")

// ErrEmptyRecipes Error returned when firmware has no recipes
var ErrEmptyRecipes = errors.New("recipes for firmware is empty")

// FirmwareRecipe ESP32 firmware is built from several binary files with their own offset.
// This struct represents the path to a binary file and its offset.
type FirmwareRecipe struct {
	Offset uint   // offset of the binary file, usually represented in hexadecimal number
	Path   string // relative path to the binary file.
}

// Firmware ESP32 firmware object
type Firmware struct {
	outputPath string
	recipes    []FirmwareRecipe
	buffer     bytes.Buffer
}

// SetOutputPath set the path to the output file of the firmware
func (f *Firmware) SetOutputPath(outputPath string) error {
	if len(outputPath) < 1 {
		return ErrEmptyOutputPath
	}

	f.outputPath = outputPath

	return nil
}

// SetRecipes set the recipes for the firmware
func (f *Firmware) SetRecipes(recipes []FirmwareRecipe) error {
	if len(recipes) < 1 {
		return ErrEmptyRecipes
	}

	f.recipes = recipes

	return nil
}

// Build build the firmware from recipes
func (f *Firmware) Build() error {
	var fatalError error
	defer func() {
		if fatalError != nil {
			os.Remove(f.outputPath)
		}
	}()

	file, err := os.Create(f.outputPath)
	defer file.Close()

	if err != nil {
		return err
	}

	var currentOffset uint
	var readLength uint
	var reader *bufio.Reader

	currentOffset = 0x1000
	for _, recipe := range f.recipes {
		if recipe.Offset < currentOffset {
			fatalError = errors.New("offset does not match")
			return fatalError
		}

		padWith(file, 0xFF, recipe.Offset-currentOffset)
		currentOffset = recipe.Offset

		recipeFile, err := os.Open(recipe.Path)
		reader = bufio.NewReader(recipeFile)
		defer recipeFile.Close()

		if err != nil {
			fatalError = err
			return err
		}

		f.buffer.Reset()
		readLength, err = writeToBuffer(&f.buffer, reader)

		if err != nil {
			fatalError = err
			return err
		}

		currentOffset += readLength
		_, err = file.Write(f.buffer.Bytes())

		if err != nil {
			fatalError = err
			return err
		}
	}

	return nil
}
