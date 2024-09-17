// 3) Dado um vetor que guarda o valor de faturamento diário de uma distribuidora, faça um programa, na linguagem que desejar, que calcule e retorne:
// • O menor valor de faturamento ocorrido em um dia do mês;
// • O maior valor de faturamento ocorrido em um dia do mês;
// • Número de dias no mês em que o valor de faturamento diário foi superior à média mensal.

// IMPORTANTE:
// a) Usar o json ou xml disponível como fonte dos dados do faturamento mensal;
// b) Podem existir dias sem faturamento, como nos finais de semana e feriados. Estes dias devem ser ignorados no cálculo da média;

// SOLUÇÂO: 
// 1. Ler o arquivo json passado como argumento
// 2. Decodificar o arquivo json para um slice de Faturamento
// 3. Calcular o menor e o maior valor de faturamento
// 4. Calcular a média mensal
// 5. Calcular o número de dias em que o valor de faturamento diário foi superior à média mensal

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type Faturamento struct {
	Dia          int     `json:"dia"`
	Valor        float64 `json:"valor"`
}

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Uso: go run ex03.go <arquivo.json>")
		return
	}
	faturamentos, err := lerFaturamentos(args[1])
	if err != nil {
		fmt.Println("Erro ao ler faturamentos:", err)
		return
	}
	if len(faturamentos) == 0 {
		fmt.Println("Nenhum dado de faturamento encontrado no arquivo.")
		return
	}

	exibirResultados(faturamentos)
}

func exibirResultados(faturamentos []Faturamento) {
	min, max := minMax(faturamentos)
	media := mediaMensal(faturamentos)
	diasAcima := diasAcimaDaMedia(faturamentos, media)
	fmt.Printf("Menor valor de faturamento: R$ %.2f\n", min)
	fmt.Printf("Maior valor de faturamento: R$ %.2f\n", max)
	fmt.Printf("Média mensal de faturamento: R$ %.2f\n", media)
	fmt.Printf("Dias em que o valor de faturamento foi superior à média: %d\n", diasAcima)
}

func lerFaturamentos(arquivo string) ([]Faturamento, error) {
	var faturamentos []Faturamento
	dados, err := ioutil.ReadFile(arquivo)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dados, &faturamentos)
	if err != nil {
		return nil, err
	}
	return faturamentos, nil
}

func minMax(faturamentos []Faturamento) (float64, float64) {
	var valores []float64
	for _, f := range faturamentos {
		if f.Valor > 0 {
			valores = append(valores, f.Valor)
		}
	}
	sort.Float64s(valores)
	return valores[0], valores[len(valores)-1]
}

func mediaMensal(faturamentos []Faturamento) float64 {
	var soma float64
	var dias int
	for _, f := range faturamentos {
		if f.Valor > 0 {
			soma += f.Valor
			dias++
		}
	}
	return soma / float64(dias)
}

func diasAcimaDaMedia(faturamentos []Faturamento, media float64) int {
	var dias int
	for _, f := range faturamentos {
		if f.Valor > media {
			dias++
		}
	}
	return dias
}