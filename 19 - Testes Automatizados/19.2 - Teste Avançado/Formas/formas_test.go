package Formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 25}
		areaEsperada := float64(250)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{25}
		areaEsperada := float64(math.Pi * 625)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

}
