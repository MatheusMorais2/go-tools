package templates

import (
	"fmt"
	"go-tools/pkg/utils"
	"log"
	"os"
	"runtime"
)

func WritePortFile(basePath, entity string) error {
	path := fmt.Sprintf("%s/internal/core/port/%s", basePath, entity)
	data := writePortData(entity)
	err := os.WriteFile(utils.GoFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writePortData(entity string) []byte {
	capitalizedEntity := utils.CapitalizeFirstLetter(entity)
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
