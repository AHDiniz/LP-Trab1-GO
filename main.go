/**
 * Linguagens de Programação - Trabalho 1
 *
 * Alan Herculano Diniz
 *
 * Solve grouping problem with the leader algorithm
 *
 * main.go: program's entry point
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
}
