package main

import (
	"errors"
	"fmt"
)

type FiguraGeometrica interface {
	Area() int
	Llenar(valores []int) error
	Nombre() string
	Perimetro() int
}

func GenerarReporte(figura FiguraGeometrica) {
	fmt.Printf("Nombre de la figura: %s\n", figura.Nombre())
	fmt.Printf("Area:                %d\n", figura.Area())
	fmt.Printf("Perimetro:           %d\n", figura.Perimetro())
}

func GenerarFiguraGeometrica(numeroLados int, valores ...int) (FiguraGeometrica, error) {
	switch numeroLados {
	case 3:
		triangulo := Triangulo{}
		triangulo.Llenar(valores)
		return &triangulo, nil
	case 4:
		rectangulo := Rectangulo{}
		rectangulo.Llenar(valores)
		return &rectangulo, nil
	default:
		return nil, errors.New("no se puede generar una figura con el numero de lados proporcionados")
	}
}

type Rectangulo struct {
	Base   int
	Altura int
}

func (r Rectangulo) Area() int {
	return r.Base * r.Altura
}

func (r *Rectangulo) Llenar(valores []int) error {
	if len(valores) < 2 {
		return errors.New("no se puede generar un rectangulo con menos de dos valores")
	}

	if valores[0] <= 0 {
		return errors.New("no se puede generar un rectangulo con una base menor o igual que 0")
	}

	if valores[1] <= 0 {
		return errors.New("no se puede generar un rectangulo con una altura menor o igual que 0")
	}

	r.Base = valores[0]
	r.Altura = valores[1]

	return nil
}

func (r Rectangulo) Nombre() string {
	return "RECTANGULO"
}

func (r Rectangulo) Perimetro() int {
	return 2*r.Base + 2*r.Altura
}

type Triangulo struct {
	Base   int
	Altura int
}

func main() {
	miFigura, err := GenerarFiguraGeometrica(4, 2, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		GenerarReporte(miFigura)
	}
}
