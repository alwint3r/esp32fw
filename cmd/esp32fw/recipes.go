package main

import (
	"bufio"
	"github.com/alwint3r/esp32fw"
	"io"
	"os"
	"path/filepath"
	"strings"
	"encoding/csv"
	"strconv"
	"errors"
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

			return splittedPartitionName[0], nil
		}

		if err == io.EOF {
			return partitionName, nil
		} else if err != nil && err != io.EOF {
			return "", err
		}
	}
}

func getFactoryAppOffset(projectDir, partitionFilename string) (uint, error) {
	partitionFile := filepath.Join(projectDir, partitionFilename+".csv")
	_, err := os.Stat(partitionFile)

	if err != nil {
		return 0, err
	}

	openedFile, err := os.Open(partitionFile)
	if err != nil {
		return 0, err
	}

	reader := csv.NewReader(bufio.NewReader(openedFile));
	reader.Comment = '#'
	reader.Comma = ','

	var found bool
	var foundOffset uint64
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}

		if (strings.TrimSpace(line[1]) == "app" || strings.TrimSpace(line[1]) == "0") && strings.TrimSpace(line[2]) == "ota_0" {
			found = true;

			foundOffset, err = strconv.ParseUint(strings.TrimSpace(line[3]), 0, 64)
			if err != nil {
				return 0, err
			}

			break;
		}
	}

	if !found {
		return 0, errors.New("offset for factory app is not found")
	}

	return uint(foundOffset), nil
}

func getIdfOnlyRecipes(projectName, projectDirectory string) ([]esp32fw.FirmwareRecipe, error) {
	recipes := make([]esp32fw.FirmwareRecipe, 3)
	partitionTableName, err := getPartitionTableName(projectDirectory)

	if err != nil {
		return recipes, err
	}

	factoryAppOffset, err := getFactoryAppOffset(projectDirectory, partitionTableName)
	if err != nil {
		factoryAppOffset = 0x10000
	}

	recipes[0] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build/bootloader/bootloader.bin"),
		Offset: 0x1000,
	}

	recipes[1] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build", partitionTableName+".bin"),
		Offset: 0x8000,
	}

	recipes[2] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build", projectName+".bin"),
		Offset: factoryAppOffset,
	}

	return recipes, nil
}

func getArduinoRecipes(projectName, arduinoDirectory, projectDirectory string) ([]esp32fw.FirmwareRecipe, error) {
	recipes := make([]esp32fw.FirmwareRecipe, 4)
	defaultFilePath := filepath.Join(projectDirectory, "build", "default.bin")
	
	var partitionName string
	_, err := os.Stat(defaultFilePath);
	if err != nil {
		partitionName, err = getPartitionTableName(projectDirectory)
		if err != nil {
			return recipes, err
		}
	} else {
		partitionName = "default"
	}

	factoryAppOffset, err := getFactoryAppOffset(projectDirectory, partitionName)
	if err != nil {
		factoryAppOffset = 0x10000
	}

	recipes[0] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build/bootloader/bootloader.bin"),
		Offset: 0x1000,
	}

	recipes[1] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build", partitionName+".bin"),
		Offset: 0x8000,
	}

	recipes[2] = esp32fw.FirmwareRecipe{
		Path: filepath.Join(projectDirectory, "components", arduinoDirectory, "tools/partitions/boot_app0.bin"),
		Offset: 0xe000,
	}

	recipes[3] = esp32fw.FirmwareRecipe{
		Path:   filepath.Join(projectDirectory, "build", projectName+".bin"),
		Offset: factoryAppOffset,
	}

	return recipes, nil
}