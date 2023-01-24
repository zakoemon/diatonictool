package data

import (
	"encoding/json"
	"log"

	"github.com/zakoemon/diatonictool/internal/jsontool"
)

type Type struct {
	Name string `json:"type"`
	Keys []Key  `json:"keys"`
}

type Key struct {
	Name      string   `json:"key"`
	Notes     []string `json:"notes"`
	SharpFlat string   `json:"#/b"`
}

type Types struct {
	Types []Type
}

func (types Types) Convert(data []byte) ([]Type, error) {
	err := json.Unmarshal(data, &types.Types)
	if err != nil {
		log.Fatal(err)
	}
	return types.Types, nil
}

func NewScaleEntity() ([]Type, error) {

	types := Types{}
	types.Types = []Type{}
	return jsontool.Read[[]Type]("config/scales.json", &types)
}
