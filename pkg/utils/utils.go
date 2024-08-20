package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func CapitalizeFirstLetter(str string) string {
	firstLetter := string(str[0])
	restOfString := strings.TrimLeft(str, firstLetter)
	capitalizedFirstLetter := strings.ToUpper(firstLetter)
	return fmt.Sprintf("%s%s", capitalizedFirstLetter, restOfString)
}

func UndoChanges(basePath string) {
	cmd := exec.Command("rm", "-r", basePath)
	cmd.Run()
}

func GoFile(fileName string) string {
	fileName = fmt.Sprintf("%s.go", fileName)
	return 	fileName
}