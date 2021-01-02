package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func AppendFile(fileName string, content string) {
	createFileIfNotExist(fileName)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", fileName)
	}
	_, err = file.WriteString(content)
	if err != nil {
		panic(fmt.Sprintf("Failed appending to file: %s", fileName))
	}
	file.Close()
}

func ReadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Failed reading content from file: %s", fileName))
	}
	return fmt.Sprint(content)
}

func WriteFile(fileName string, content []byte) {
	createFileIfNotExist(fileName)
	err := ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed writing content to file: %s", fileName))
	}
}

func createFileIfNotExist(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		emptyFile, err := os.Create(fileName)
		if err != nil {
			panic(fmt.Sprintf("Failed to create empty file: %s", fileName))
		}
		emptyFile.Close()
	}
}

func CreateDir(dirName string) {
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		return
	}
	err := os.Mkdir(dirName, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed creating dir: %s", dirName))
	}
}
