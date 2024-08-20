package templates

import (
	"fmt"
	"go-tools/pkg/utils"
	"log"
	"os"
	"runtime"
)

// Eu quero montar uma serie de rotas com operacoes CRUD das entidades
// que eu passo no comando de run
// Rotas: post, get by id, update, delete

func WriteMainHttpServerFile(basePath string, entities []string) error {
	path := fmt.Sprintf("%s/internal/adapters/http/server", basePath)
	data := writeMainHttpServerData(entities)
	err := os.WriteFile(utils.GoFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}




func writeMainHttpServerData(entities []string) []byte {
	fileString := `package httpAdapter

import (
	"fmt"
	"net/http"
)

func Server()  {
	router := http.NewServeMux()
`

	for i := 0; i < len(entities); i++ {
		capitalizedEntity := utils.CapitalizeFirstLetter(entities[i])
		fileString = fileString + fmt.Sprintf(`
	%sRouter(router)`, capitalizedEntity)
	}
	fileString = fileString + fmt.Sprintf(`

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router,
	}
		
	fmt.Println("Server listening on localhost:8080")
	server.ListenAndServe()
}`)
	return []byte(fileString)
}

func WriteHttpAdapterFile(basePath string, entity string) error {
	path := fmt.Sprintf("%s/internal/adapters/http/%s", basePath, entity)
	data := writeHttpAdapterData(entity)
	err := os.WriteFile(utils.GoFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writeHttpAdapterData(entity string) []byte {
	capitalizedEntity := utils.CapitalizeFirstLetter(entity)
	fileString := fmt.Sprintf(`package httpAdapter

import (
	"fmt"
	"net/http"
)

func %sRouter(router *http.ServeMux) *http.ServeMux {
	router.HandleFunc("POST /%s", handlePost%s)
	router.HandleFunc("GET /%s/{id}", handleGetById%s)
	router.HandleFunc("UPDATE /%s", handleUpdate%s)
	router.HandleFunc("DELETE /%s", handleDelete%s)
	
	return router
}

func handlePost%s(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Print(ctx)
	w.Write([]byte("POST %s"))
}

func handleGetById%s(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Print(ctx)
	id := r.PathValue("id")
	fmt.Println(id)

	w.Write([]byte("get %s"))
}

func handleUpdate%s(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Print(ctx)
	w.Write([]byte("UPDATE %s"))
}

func handleDelete%s(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Print(ctx)
	w.Write([]byte("DELETE %s"))
}
	`,capitalizedEntity, 
	entity, capitalizedEntity,
	entity, capitalizedEntity,
	entity, capitalizedEntity,
	entity, capitalizedEntity,
	 capitalizedEntity, capitalizedEntity, capitalizedEntity, capitalizedEntity, 
	capitalizedEntity, capitalizedEntity, capitalizedEntity, capitalizedEntity)
	byteArray := []byte(fileString)
	return byteArray
}
