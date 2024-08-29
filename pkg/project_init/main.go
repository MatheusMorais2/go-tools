package main

import (
	"fmt"
	"go-tools/pkg/project_init/templates"
	"go-tools/pkg/utils"
	"log"
	"os"
	"runtime"
)

/*

	Ordem dos parametros do comando de buildar:
	1 - project name
	2 - path
	3 +: entidades

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
	os.Mkdir(basePath, 0750)

	entities, err := getEntities(os.Args)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		utils.UndoChanges(basePath)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	err = makeHexagonalDirectories(basePath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		utils.UndoChanges(basePath)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	fmt.Printf("entities: %+v\n", entities)
	err = writeCoreFiles(projectName, basePath, entities)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		utils.UndoChanges(basePath)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	err = writeAdapters(basePath, entities)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		utils.UndoChanges(basePath)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	err = writePostgresAdapter(basePath, entities)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

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

	err = os.Mkdir(fmt.Sprintf("%s/internal/adapters/storage/sql", basePath), 0750)
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

func writeCoreFiles(projectName, basePath string, entities []string) error {
	for i := 0; i < len(entities); i++ {
		err := templates.WriteDomainFile(basePath, entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}

		err = templates.WritePortFile(projectName, basePath, entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}

		err = templates.WriteServiceFile(projectName ,basePath, entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}
	}

	return nil
}

//TODO
func writeAdapters(basePath string, entities []string) error {
	err := writeHttpAdapter(basePath, entities)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	fmt.Println("fuck git")
	return nil
}

func writeHttpAdapter(basePath string, entities []string) error {
	err := templates.WriteMainHttpServerFile(basePath, entities)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}

	for i := 0; i < len(entities); i++ {
		templates.WriteHttpAdapterFile(basePath, entities[i])
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
		}

	}	
	return nil
}

//TODO
func writePostgresAdapter(basePath string, entities []string) error {
	err := templates.WriteMainPostgresRepositoryFile(basePath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

//TODO
func writeRedisAdapter(basePath string, entities []string) error {
	return nil
}

//TODO
func writeMain(basePath string) error {
	return nil
}