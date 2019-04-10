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
	"flag"
)

func main() {

	// Getting the input files location:
	pointsFile := flag.String("points_file", "entrada.txt", "Arquivo com as coordenadas dos pontos.")
	distFile := flag.String("dist_file", "distancia.txt", "Arquivo com a distância limite entre pontos.")

	points, dist := readFile(*pointsFile, *distFile) // Reading the files properlly

	fmt.Println(points, dist)

	// 2 - Criar as estruturas de dados que vão representar os dados lidos
	// 3 - Agrupar os pontos usando o algoritmo de líder
	// 4 - Criar e escrever nos arquivos de saída os resultados X
}
