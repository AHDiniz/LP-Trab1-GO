/**
 * Linguagens de Programação - Trabalho 1
 *
 * Alan Herculano Diniz
 *
 * Resolver problema de agrupamento com o algoritmo dado em aula
 *
 * main.go: ponto de entrada do programa
 */

package main

import (
	"fmt"
)

func main() {
	// 1 - Ler os arquivos de entrada no formato correto
	// 2 - Criar as estruturas de dados que vão representar os dados lidos
	// 3 - Agrupar os pontos usando o algoritmo de líder
	// 4 - Criar e escrever nos arquivos de saída os resultados

	// Testando função de leitura de arquivo
	pointsStr := readFile("entrada.txt") // Lendo o arquivo com os dados das coordenadas dos pontos
	fmt.Println(readFile("entrada.txt"))
}
