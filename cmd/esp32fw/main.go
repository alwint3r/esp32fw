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
	recipes, err := getIdfOnlyRecipes(projectName, *projectDir)

	if err != nil {
		fmt.Println("Failed getting recipes:", err)
		os.Exit(1)
	}

	firmware := esp32fw.Firmware{}
	firmware.SetOutputPath(filepath.Join(cwd, "firmware.bin"))
	firmware.SetRecipes(recipes)

	err = firmware.Build()
	if err != nil {
		fmt.Println("Failed building firmware:", err)
		os.Exit(1)
	}
}
