package main

import (
	"bufio"
	"github.com/alwint3r/esp32fw"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getPartitionTableName(projectDirectory string) (string, error) {
	sdkconfig := filepath.Join(projectDirectory, "sdkconfig")
	file, err := os.Open(sdkconfig)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(file)
	var partitionName string

	for {
		line, _, err := reader.ReadLine()
		strline := strings.TrimSpace(string(line))

		splitted := strings.Split(strline, "=")
		if splitted[0] == "CONFIG_PARTITION_TABLE_FILENAME" {
			partitionName = strings.Replace(splitted[1], "\"", "", -1)
			splittedPartitionName := strings.Split(partitionName, ".")

			return splittedPartitionName[0]+".bin", nil
		}

		if err == io.EOF {
			return partitionName, nil
		} else if err != nil && err != io.EOF {
			return "", err
		}
	}
}

func getIdfOnlyRecipes(projectName, projectDirectory string) ([]esp32fw.FirmwareRecipe, error) {
	recipes := make([]esp32fw.FirmwareRecipe, 3)
	partitionTableName, err := getPartitionTableName(projectDirectory)

	if err != nil {
		return recipes, err
	}

	recipes[0] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build/bootloader/bootloader.bin"),
		Offset: 0x1000,
	}

	recipes[1] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build", partitionTableName),
		Offset: 0x8000,
	}

	recipes[2] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build", projectName+".bin"),
		Offset: 0x10000,
	}

	return recipes, nil
}
