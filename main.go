package goarea

import "math"

// PI é uma proporção numérica definida pela relação entre o perímetro de uma
// circunferência e seu diâmetro.
const PI = 3.1416

// Circ é responsável por calcular a área da circunferência.
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

// Rect é reponsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Função privada NO PACOTE!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
