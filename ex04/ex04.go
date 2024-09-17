// 4) Dado o valor de faturamento mensal de uma distribuidora, detalhado por estado:
// • SP – R$67.836,43
// • RJ – R$36.678,66
// • MG – R$29.229,88
// • ES – R$27.165,48
// • Outros – R$19.849,53

// Escreva um programa na linguagem que desejar onde calcule o percentual de representação que cada estado teve dentro do valor total mensal da distribuidora.

package main

import (
	"fmt"
)

func main() {
	faturamento := map[string]float64{
		"SP": 67836.43,
		"RJ": 36678.66,
		"MG": 29229.88,
		"ES": 27165.48,
		"Outros": 19849.53,
	}
	total := calcular_percentual(faturamento)
	print_percentual(faturamento, total)
}

func calcular_percentual(faturamento map[string]float64) float64 {
	total := 0.0
	for _, valor := range faturamento {
		total += valor
	}
	return (total)
}

func print_percentual(faturamento map[string]float64, total float64) {
	fmt.Printf("Total: R$ %.2f\n", total)
	fmt.Println("Percentual de representação:")
	for estado, valor := range faturamento {
		percentual := (valor / total) * 100
		fmt.Printf("%s: %.2f%%\n", estado, percentual)
	}
}