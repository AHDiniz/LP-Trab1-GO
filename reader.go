/**
 * Linguagens de Programação - Trabalho 1
 *
 * Alan Herculano Diniz
 *
 * Solve groupping problem with the leader algorithm
 *
 * reader.go: I/O file manipulation
 *
 * **************************************************************************************
 *
 * Things to do:
 *
 * 1 - Create function to parse string into a slice of points (wich are slices of floats)
 */

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"bufio"
	"fmt"
	"strings"
)

/**
 * Reading the input file
 *
 * This reads the file to get the point coordinates and the limit distance to create the groups
 *
 * Input: the files that must be read
 *
 * Output: a the point slice and the maximum distance
 */
func readFile(pointsLoc, distLoc string) ([][]float64, float64) {

	// Opening the points file:
	pointsFile, err := os.Open(pointsLoc) // Opening the file
	defer pointsFile.Close()

	if err != nil {
		panic(err)
	}

	// Opening the distance file:
	distFile, err := os.Open(distLoc)
	defer distFile.Close()

	if err != nil {
		panic(err)
	}

	// Reading the limit distance:
	var dist float64
	fmt.Fscanf(distFile, "%f", &dist)

	// Getting the file's lines and the number of them:
	lines, linesNum := getFileLines(pointsFile)

	// Parsing the string slice to a point slice:
	points := parseInputString(lines, linesNum)

	return points, dist
}

/**
 * Getting a string slice and returning the points
 *
 * Input: string slice with the input file lines and the amount of lines in the file
 *
 * Output: the float matrix containing the points' coordinates
 */
func parseInputString(input []string, lines int) [][]float64 {

	// Determining the amount of coordinates of each point:
	dimension := len(strings.Fields(input[0]))

	// Creating the point slice:
	var points [][]float64 = make([][]float64, lines)
	for i := 0; i < lines; i++ {
		points[i] = make([]float64, dimension)
	}

	// Getting the points from the string slice:
	for i := 0; i < lines; i++ {

		coordStrs := strings.Fields(input[i]) // Getting a slice with each coordinate in a string

		// Getting the corrdinates in the string:
		for j := 0; j < dimension; j++ {

			fmt.Sscanf(coordStrs[j], "%f", &points[i][j])
		}
	}

	return points
}

/**
 * Getting a file and returning it's lines
 *
 * Input: the file that must be read
 *
 * Outputs: the string slice with the lines and the amount of lines in the file
*/
func getFileLines(input *os.File) ([]string, int) {

	// Determining how many points where given in the file:
	scanner := bufio.NewScanner(input)
	var lines []string
	lineNum := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineNum++
	}

	return lines, lineNum
}

/**
 * Impressão dos resultados do algoritmo
 *
 * A soma euclidiana das distâncias entre os pontos será impressa num
 * arquivo result.txt por meio desta função, assim como os grupos, que
 * serão impressos no arquivo saida.txt
 *
 * Entradas: resultado da soma das distâncias euclidianas e slice com os
 * pontos agrupados
 */
func printResults(sse int, groups [][]int) {

	content := strconv.Itoa(sse) // Convertendo a soma para uma string

	err := ioutil.WriteFile("result.txt", []byte(content), os.ModeAppend) // Imprimindo a soma no arquivo devido

	// Detecção de erro na abertura e escrita do arquivo:
	if err != nil {
		log.Fatalln("The result.txt file can't be created for some reason...")
	}

	// Criação da string com os grupos resultantes do programa:
	var groupsStr string
	for i := 0; i < len(groups); i++ {
		for j := 0; j < len(groups[i]); j++ {

			groupsStr += strconv.Itoa(groups[i][j]) + " "

		}
		groupsStr += "\n\n"
	}

	err = ioutil.WriteFile("saida.txt", []byte(groupsStr), os.ModeAppend) // Imprimindo o arquivo com os grupos criados no algoritmo

	// Detecção de erro na abertura e escrita do arquivo:
	if err != nil {
		log.Fatalln("The saida.txt file can't be created for some reason...")
	}
}
