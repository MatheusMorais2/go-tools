package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main()  {
	if (len(os.Args) > 1) {
		fmt.Printf("First arg passed: %s\n", os.Args[1])
	}

	scanner := bufio.NewScanner(os.Stdin)
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter project name: ")
	scanner.Scan()
	projectName := scanner.Text()
	fmt.Print("Enter path: ")
	scanner.Scan()
	path := scanner.Text()

	err = os.Mkdir("pkg", 0750)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(fmt.Sprintf("pkg/%s", path), 0750)
	if err != nil {
		log.Fatal(err)
	}

 	data := []byte{1, 2, 3}
	err = os.WriteFile(goFile(projectName), data, 0600)
	if err != nil {
		log.Fatal(err)
	} 

	fmt.Printf("Creating project %s at %s\n", projectName, path)
}

func goFile(fileName string) string {
	fileName = fmt.Sprintf("%s.go", fileName)
	return 	fileName
}
