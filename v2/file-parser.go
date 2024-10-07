package main

import "os"

func FileParser(filepath string, entity string) error {
	fileBytes, err := os.ReadFile(filepath)
	err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
}
	
