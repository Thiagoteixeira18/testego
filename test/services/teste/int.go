package services

import (
	"fmt"
	"math"
)

func SomeValue() string {
	return "expect value"
}

func Principal() {
	SomeValue := SomeValue()

	addition := Sum(3, 2)

	Multiplicacao := Multpy(5, -6)

	fmt.Printf("Some Value: %s, Addition: %d, multiplicacao: %d/n",
		SomeValue, addition, Multiplicacao)
}

func Sum(a int, b int) int {
	return a + b
}

func Multpy(a int, b int) int {
	return a * b
}

func multi(d, t int) int {
	return d * t
}

const quantidadeRepeticoes = 5

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}

	return somas
}

type Forma interface {
	Area() float64
}

// Retangulo tem as dimensões de um retângulo
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Area retorna a área de um retângulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Perimetro retorna o perímetro de um retângulo
func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

// Circulo representa um círculo.
type Circulo struct {
	Raio float64
}

// Area retorna a área de um círculo
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

// Triangulo representa as dimensões de um triângulo
type Triangulo struct {
	Base   float64
	Altura float64
}

// Area retorna a área de um triângulo
func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}
