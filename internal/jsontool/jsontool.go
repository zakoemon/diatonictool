package jsontool

import (
	"encoding/json"
	"log"
	"os"
)

func Read[T any](path string, t T) (T, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &t)
	if err != nil {
		log.Fatal(err)
	}

	return t, nil
}
