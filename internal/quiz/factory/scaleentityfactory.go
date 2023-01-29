package factory

import (
	"github.com/zakoemon/diatonictool/internal/data"
)

func randomKeyGenerator(scaleType data.ScaleType) func() data.Key {
	cache, _ := data.NewScaleEntity()
	t := int(scaleType)
	keys := cache[t].Keys

	return func() data.Key {
		num := genRand(len(keys))
		return keys[num]
	}
}
