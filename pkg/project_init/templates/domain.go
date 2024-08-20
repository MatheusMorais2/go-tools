package templates

import (
	"fmt"
	"go-tools/pkg/utils"
	"log"
	"os"
	"runtime"
)

func WriteDomainFile(basePath, entity string) error {
	path := fmt.Sprintf("%s/internal/core/domain/%s", basePath, entity)
	data := writeDomainData(entity)
	err := os.WriteFile(utils.GoFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writeDomainData(entity string) []byte {
	capitalizedEntity := utils.CapitalizeFirstLetter(entity)
	fileString := fmt.Sprintf(`package domain

type %s struct {}`, capitalizedEntity)
	byteArray := []byte(fileString)
	return byteArray
}