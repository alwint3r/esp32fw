package esp32fw

import "testing"

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
