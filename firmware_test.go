package esp32fw

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInvalidValuesHandling(t *testing.T) {
	firmware := Firmware{}

	err := firmware.SetOutputPath("")
	if err == nil {
		t.Error("Should return error if the output path is an empty string")
	}

	if err != nil && err != ErrEmptyOutputPath {
		t.Error("Should return ErrEmptyOutputPath if path is empty")
	}

	err = firmware.SetRecipes([]FirmwareRecipe{})
	if err == nil {
		t.Error("Should return error if the recipes are empty")
	}

	if err != nil && err != ErrEmptyRecipes {
		t.Error("Should return ErrEmptyRecipes")
	}
}

func TestFirmwareCreation(t *testing.T) {
	firmware := Firmware{}
	workingDir, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	outputPath := filepath.Join(workingDir, "test_files/output.bin")
	recipes := []FirmwareRecipe{
		FirmwareRecipe{
			Offset: 0x1000,
			Path:   filepath.Join(workingDir, "test_files/bootloader.bin"),
		},
		FirmwareRecipe{
			Offset: 0x8000,
			Path:   filepath.Join(workingDir, "test_files/partitions.bin"),
		},
		FirmwareRecipe{
			Offset: 0x10000,
			Path:   filepath.Join(workingDir, "test_files/main.bin"),
		},
	}

	err = firmware.SetOutputPath(outputPath)
	if err != nil {
		t.Error(err)
	}

	err = firmware.SetRecipes(recipes)
	if err != nil {
		t.Error(err)
	}

	err = firmware.Build()
	if err != nil {
		t.Error(err)
	}

	_, err = os.Open(outputPath)
	if err != nil {
		t.Error("Firmware is not created!")
	}
}
