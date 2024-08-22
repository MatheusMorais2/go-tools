package templates

import (
	"fmt"
	"go-tools/pkg/utils"
	"log"
	"os"
	"runtime"
)

func WriteServiceFile(projectName, basePath, entity string) error {
	path := fmt.Sprintf("%s/internal/core/service/%s", basePath, entity)
	data := writeServiceData(projectName, entity)
	err := os.WriteFile(utils.GoFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writeServiceData(projectName string,entity string) []byte {
	capitalizedEntity := utils.CapitalizeFirstLetter(entity)
	fileString := fmt.Sprintf(`package service
import (
	"context"
	"go-tools/%s/internal/core/domain"
)

type %sService struct {}

func New%sService() *%sService {
	return &%sService{
	}
}
	
func (s *%sService) Create(ctx context.Context, %s *domain.%s) (*domain.%s, error) {
	return nil, nil
}

func (s *%sService) Get(ctx context.Context, %s *domain.%s) (*domain.%s, error) {
	return nil, nil
}

func (s *%sService) Update(ctx context.Context, %s *domain.%s) (*domain.%s, error) {
	return nil, nil
}

func (s *%sService) Delete(ctx context.Context, %s *domain.%s) (*domain.%s, error) {
	return nil, nil
}`, projectName, capitalizedEntity, 
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