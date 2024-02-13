package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/scripts/openapi-spec-generator/info"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/scripts/openapi-spec-generator/path"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/scripts/openapi-spec-generator/security"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/scripts/openapi-spec-generator/server"
	"github.com/swaggest/openapi-go/openapi3"
)

const filePath string = "./api/openapi-spec.yaml"

func main() {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}

	info.AddInfo(&reflector)
	server.AddServers(&reflector)

	path.AddPlayerOperations(&reflector)
	path.AddGameOperations(&reflector)

	security.AddSecurity(&reflector)

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(schema)

	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Successfully generated openapi spec, saved to file: %s", filePath)
}
