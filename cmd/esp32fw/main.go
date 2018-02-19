package main

import (
	"flag"
	"fmt"
	"github.com/alwint3r/esp32fw"
	"os"
	"path/filepath"
)

func main() {
	projectDir := flag.String("project", "", "Path to ESP-IDF project")
	isUsingArduino := flag.Bool("use-arduino", false, "Set this flag if you're using arduino as component to ESP-IDF project")
	arduinoDirectoryName := flag.String("arduino-dir", "arduino-esp32", "Directory name of the arduino ESP32 component")
	outputPathOption := flag.String("output-path", "", "Path to firmware output")

	flag.Parse()

	if len(*projectDir) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory")
		os.Exit(1)
	}

	if *projectDir == "." {
		*projectDir = cwd
	}

	projectName := filepath.Base(*projectDir)

	var actualOutputPath string
	if len(*outputPathOption) < 1 || *outputPathOption == "." {
		actualOutputPath = filepath.Join(cwd, projectName+".bin")
	} else {
		actualOutputPath = *outputPathOption
	}

	var recipes []esp32fw.FirmwareRecipe
	if *isUsingArduino == true {
		fmt.Println("This project is using Arduino as component")
		recipes, err = getArduinoRecipes(projectName, *arduinoDirectoryName, *projectDir)
	} else {
		recipes, err = getIdfOnlyRecipes(projectName, *projectDir)
	}

	if err != nil {
		fmt.Println("Failed getting recipes:", err)
		os.Exit(1)
	}

	firmware := esp32fw.Firmware{}
	firmware.SetOutputPath(actualOutputPath)
	firmware.SetRecipes(recipes)

	err = firmware.Build()
	if err != nil {
		fmt.Println("Failed building firmware:", err)
		os.Exit(1)
	}

	fmt.Printf("Saved firmware to %s\n", actualOutputPath)
}
