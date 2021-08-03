package main

import (
	"errors"
	"fmt"
	"math"
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
	case 0:
		circulo := Circulo{0}
		circulo.Llenar(valores)
		return &circulo, nil
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
	Base    int
	Altura  int
	LadoUno int
	LadoDos int
}

func (t Triangulo) Area() int {
	return t.Base * t.Altura / 2
}

func (t *Triangulo) Llenar(valores []int) error {
	if len(valores) != 4 {
		return errors.New("No se puede generar un triangulo con menos de cuatro valores")
	}

	if valores[0] <= 0 {
		return errors.New("No se puede generar un triangulo con una base menor o igual que cero")
	}

	if valores[1] <= 0 {
		return errors.New("No se puede generar un triangulo si su altura es menor o igual que cero")
	}

	if valores[2] <= 0 {
		return errors.New("No se puede generar un triangulo si uno de sus lados es menor o igual que cero")
	}

	if valores[3] <= 0 {
		return errors.New("No se puede generar un triangulo si uno de sus lados es menor o igual que cero")
	}

	t.Base = valores[0]
	t.Altura = valores[1]
	t.LadoUno = valores[2]
	t.LadoDos = valores[3]

	return nil
}

func (t Triangulo) Nombre() string {
	return "Triangulo"
}

func (t Triangulo) Perimetro() int {
	return t.Base + t.LadoUno + t.LadoDos
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() int {
	return int(math.Pi * c.Radio * c.Radio)
}

func (c *Circulo) Llenar(valores []int) error {

	if len(valores) != 1 {
		errors.New("No se puede generar ")
	}
	if valores[0] <= 0 {
		errors.New("No se puede generar un circulo con un radio menor que cero")
	}

	c.Radio = float64(valores[0])
	return nil
}

func (c Circulo) Nombre() string {
	return "Circulo"
}

func (c Circulo) Perimetro() int {
	return int(2 * math.Pi * c.Radio)
}

func main() {
	miFigura, err := GenerarFiguraGeometrica(0, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		GenerarReporte(miFigura)
	}
}
