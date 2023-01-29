package jsontool

import (
	"log"
	"os"
)

type Convert[T any] interface {
	Convert(data []byte) (T, error)
}

func Read[T any](path string, converter Convert[T]) (T, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var entity T
	entity, err = converter.Convert(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return entity, nil
}
