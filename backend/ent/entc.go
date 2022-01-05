//go:build ignore
// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/masseelch/elk"
)

func main() {
	ex, err := elk.NewExtension(
		elk.GenerateSpec("openapi.json"),
		elk.GenerateHandlers(),
	)
	if err != nil {
		log.Fatalf("creating elk extension: %v", err)
	}
	generateClient()
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func generateClient() {
	swagger, err := util.LoadSwagger("./openapi.json")
	if err != nil {
		log.Fatalf("Failed to load swagger %v", err)
	}

	generated, err := codegen.Generate(swagger, "stub", codegen.Options{
		GenerateClient: true,
		GenerateTypes:  true,
		AliasTypes:     true,
	})
	if err != nil {
		log.Fatalf("generaring client failed %s", err)
	}

	dir := filepath.Join(".", "stub")
	stub := filepath.Join(".", "stub", "http.go")
	perm := os.FileMode(0777)
	if err := os.MkdirAll(dir, perm); err != nil {
		log.Fatalf("error creating dir: %s", err)
	}

	if err := ioutil.WriteFile(stub, []byte(generated), perm); err != nil {
		log.Fatalf("error writing generated code to file: %s", err)
	}
}
