package main

import (
	"fmt"
	"log"
	"os"
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

func main()  {

	// 1 - tenho que definir o core
	entities, err := getEntities(os.Args)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	// 2 - criar a estrutura de pastas
	err = makeHexagonalDirectories()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	// 3 - Criar os arquivos do core
	fmt.Printf("entities: %+v\n", entities)
	err = writeCoreFiles(entities)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	// 4 - Criar o template dos adapters

}

func goFile(fileName string) string {
	fileName = fmt.Sprintf("%s.go", fileName)
	return 	fileName
}

func makeHexagonalDirectories() error {
	err := os.Mkdir("internal", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/adapters", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/adapters/http", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/adapters/storage", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/core", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/core/domain", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/core/port", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/core/service", 0750)
	if err != nil {
		return err
	}

	err = os.Mkdir("internal/core/util", 0750)
	if err != nil {
		return err
	}

/* 	err = os.Mkdir("cmd", 0750)
	if err != nil {
		return err
	} */

	return nil
}

func getEntities(args []string) ([]string, error) {
	fmt.Println(args)
	if (len(args) <= 1) {
		return nil, fmt.Errorf("entities not passed in arguments")
	}
	entities := make([]string, 0)
	for i := 1; i < len(args); i++ {
		entities = append(entities, args[i])
	}
	return entities, nil
}

func writeCoreFiles(entities []string) error {
	for i := 0; i < len(entities); i++ {
		err := writeDomainFile(entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}

		err = writePortFile(entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}

		err = writeServiceFile(entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}
	}

	return nil
}

func writeDomainFile(entity string) error {
	path := fmt.Sprintf("internal/core/domain/%s", entity)
	data := writeDomainData(entity)
	err := os.WriteFile(goFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writePortFile(entity string) error {
	path := fmt.Sprintf("internal/core/port/%s", entity)
	data := writePortData(entity)
	err := os.WriteFile(goFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}
 
func writeServiceFile(entity string) error {
	path := fmt.Sprintf("internal/core/service/%s", entity)
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
