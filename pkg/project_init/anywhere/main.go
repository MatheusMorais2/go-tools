package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

/*
	Tenho 2 opcoes:
	1 - Iniciar o projeto de qualquer lugar:
		1.1 - Tenho que passar o endereço de onde eu quero usar
		1.2 - Tenho que passar o nome das entidades
	2 - Importar o package e iniciar na pasta que eu to usando:
		1.1 - tenho que tornar meu module disponivel e usavel para outros
		1.2 - Tenho que passar o nome das entidades

*/

// Preciso fazer isso uma transacao
func main()  {
	basePath := getBasePath(os.Args)

	projectName, err := getProjectName(os.Args)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	basePath = fmt.Sprintf("%s/%s", basePath, projectName)
	fmt.Printf("basePath: %+v\n", basePath)
	os.Mkdir(basePath, 0750)

	// 1 - tenho que definir o core
	entities, err := getEntities(os.Args)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		//undoChanges(basePath)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	// 2 - criar a estrutura de pastas
	err = makeHexagonalDirectories(basePath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		//undoChanges(basePath)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	// 3 - Criar os arquivos do core
	fmt.Printf("entities: %+v\n", entities)
	err = writeCoreFiles(basePath, entities)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		//undoChanges(basePath)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	undoChanges(basePath)
	// 4 - Criar o template dos adapters

}

func goFile(fileName string) string {
	fileName = fmt.Sprintf("%s.go", fileName)
	return 	fileName
}

func getProjectName(args []string) (string, error) {
	var projectName string
	if len(args) <= 1 {
		return "", fmt.Errorf("project name was not passed")
	} else {
		projectName = args[1]
	}
	return projectName, nil
} 

func getBasePath(args []string) string {
	var basePath string
	fmt.Println(args)
	if len(args) <= 2 {
		basePath = "."
	} else {
		basePath = args[2]
	}
	return basePath
}

func makeHexagonalDirectories(basePath string) error {
	err := os.Mkdir(fmt.Sprintf("%s/internal", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/adapters", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/adapters/http", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/adapters/storage", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/core", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/core/domain", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/core/port", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/core/service", basePath), 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal/core/util", basePath), 0750)
	if err != nil {
		return err
	}

/* 	err = os.Mkdir(fmt.Sprintf("%s/cmd", basePath), 0750)
	if err != nil {
		return err
	} */

	return nil
}

func getEntities(args []string) ([]string, error) {
	if (len(args) <= 3) {
		return nil, fmt.Errorf("entities not passed in arguments")
	}
	entities := make([]string, 0)
	for i := 3; i < len(args); i++ {
		entities = append(entities, args[i])
	}
	return entities, nil
}

func writeCoreFiles(basePath string, entities []string) error {
	for i := 0; i < len(entities); i++ {
		err := writeDomainFile(basePath, entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}

		err = writePortFile(basePath, entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}

		err = writeServiceFile(basePath, entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}
	}

	return nil
}

func writeDomainFile(basePath, entity string) error {
	path := fmt.Sprintf("%s/internal/core/domain/%s", basePath, entity)
	data := writeDomainData(entity)
	err := os.WriteFile(goFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writePortFile(basePath, entity string) error {
	path := fmt.Sprintf("%s/internal/core/port/%s", basePath, entity)
	data := writePortData(entity)
	err := os.WriteFile(goFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}
 
func writeServiceFile(basePath, entity string) error {
	path := fmt.Sprintf("%s/internal/core/service/%s", basePath, entity)
	data := writeServiceData(entity)
	err := os.WriteFile(goFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writeServiceData(entity string) []byte {
	capitalizedEntity := capitalizeFirstLetter(entity)
	fileString := fmt.Sprintf(`package service
import "context"

type %sService struct {}

func New%sService() *%sService {
	return &%sService{
	}
}
	
func (s *%sService) Create(ctx context.Context, %s *domain.%s) (*%s, error) {
}

func (s *%sService) Get(ctx context.Context, %s *domain.%s) (*%s, error) {
}

func (s *%sService) Update(ctx context.Context, %s *domain.%s) (*%s, error) {
}

func (s *%sService) Delete(ctx context.Context, %s *domain.%s) (*%s, error) {
}`, capitalizedEntity, 
	capitalizedEntity, capitalizedEntity,
	capitalizedEntity,
	capitalizedEntity, entity, capitalizedEntity, capitalizedEntity,
	capitalizedEntity, entity, capitalizedEntity, capitalizedEntity,
	capitalizedEntity, entity, capitalizedEntity, capitalizedEntity,
	capitalizedEntity, entity, capitalizedEntity, capitalizedEntity,
)
	byteArray := []byte(fileString)
	return byteArray
}

func writePortData(entity string) []byte {
	capitalizedEntity := capitalizeFirstLetter(entity)
	fileString := fmt.Sprintf(`package port
import "context"

type %sRepository interface {
	Create(ctx context.Context, %s *domain.%s) (*domain.%s, error)
	Update(ctx context.Context, %s *domain.%s) (*domain.%s, error)
	Delete(ctx context.Context, %s *domain.%s) (*domain.%s, error)
	Get(ctx context.Context, %s *domain.%s) (*domain.%s, error)
}

type %sService interface {
	Create(ctx context.Context, %s *domain.%s) (*domain.%s, error)
	Update(ctx context.Context, %s *domain.%s) (*domain.%s, error)
	Delete(ctx context.Context, %s *domain.%s) (*domain.%s, error)
	Get(ctx context.Context, %s *domain.%s) (*domain.%s, error)
}`, capitalizedEntity, 
	entity, capitalizedEntity, capitalizedEntity,
	entity, capitalizedEntity, capitalizedEntity,
	entity, capitalizedEntity, capitalizedEntity,
	entity, capitalizedEntity, capitalizedEntity,
	capitalizedEntity,
	entity, capitalizedEntity, capitalizedEntity,
	entity, capitalizedEntity, capitalizedEntity,
	entity, capitalizedEntity, capitalizedEntity,
	entity, capitalizedEntity, capitalizedEntity,)
	byteArray := []byte(fileString)
	return byteArray
}

func writeDomainData(entity string) []byte {
	capitalizedEntity := capitalizeFirstLetter(entity)
	fileString := fmt.Sprintf(`package domain

type %s struct {}`, capitalizedEntity)
	byteArray := []byte(fileString)
	return byteArray
}

func capitalizeFirstLetter(str string) string {
	firstLetter := string(str[0])
	restOfString := strings.TrimLeft(str, firstLetter)
	capitalizedFirstLetter := strings.ToUpper(firstLetter)
	return fmt.Sprintf("%s%s", capitalizedFirstLetter, restOfString)
}

func undoChanges(basePath string) {
	cmd := exec.Command("rm", "-r", basePath)
	cmd.Run()
}
