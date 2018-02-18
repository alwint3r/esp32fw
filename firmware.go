package esp32fw

import (
	"errors"
	"bytes"
)

// ErrEmptyOutputPath Error returned when output path to a firmware is empty
var ErrEmptyOutputPath = errors.New("output path is empty")

// ErrEmptyRecipes Error returned when firmware has no recipes
var ErrEmptyRecipes = errors.New("recipes for firmware is empty")

// FirmwareRecipe ESP32 firmware is built from several binary files with their own offset.
// This struct represents the path to a binary file and its offset.
type FirmwareRecipe struct {
	Offset uint32 // offset of the binary file, usually represented in hexadecimal number
	Path string // relative path to the binary file.
}

// Firmware ESP32 firmware object
type Firmware struct {
	outputPath string
	recipes []FirmwareRecipe
	buffer bytes.Buffer
}

// SetOutputPath set the path to the output file of the firmware
func (f *Firmware) SetOutputPath(outputPath string) error {
	if len(outputPath) < 1 {
		return ErrEmptyOutputPath
	}

	return nil
}

// SetRecipes set the recipes for the firmware
func (f *Firmware) SetRecipes(recipes []FirmwareRecipe) error {
	if len(recipes) < 1 {
		return ErrEmptyRecipes
	}
	return nil
}

// Build build the firmware from recipes
func (f *Firmware) Build() error {
	return nil
}