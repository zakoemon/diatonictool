package data

import (
	"github.com/zakoemon/diatonictool/internal/jsontool"
)

type ScaleType int

const (
	Major ScaleType = 0
	Minor ScaleType = 1
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

func NewScaleEntity() ([]Type, error) {

	types := []Type{}
	return jsontool.Read("config/scales.json", types)
}
