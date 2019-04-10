/**
 * Linguagens de Programação - Trabalho 1
 *
 * Alan Herculano Diniz
 *
 * Solve grouping problem with the leader algorithm
 *
 * reader.go: I/O file manipulation
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
	var lines []string // Slice where each position is a line in the file
	lineNum := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Appending the read line in the slice
		lineNum++
	}

	return lines, lineNum
}

/**
 * Printing the algorithm's results
 *
 * The euclidian sum of distances between the groups will be printed in
 * a result.txt file, while the groups will be printed in a saida.txt file.
 *
 * Inputs: the euclidian sum of distances and the point groups
 */
func printResults(sse int, groups [][]int) {

	content := strconv.Itoa(sse) // Parsing the sum to a string

	err := ioutil.WriteFile("result.txt", []byte(content), os.ModeAppend) // Printing the sum

	// Error detection in the file operation:
	if err != nil {
		log.Fatalln("The result.txt file can't be created for some reason...")
	}

	// Creating the string with the point groups:
	var groupsStr string
	for i := 0; i < len(groups); i++ {
		for j := 0; j < len(groups[i]); j++ {

			groupsStr += strconv.Itoa(groups[i][j]) + " "

		}
		groupsStr += "\n\n"
	}

	err = ioutil.WriteFile("saida.txt", []byte(groupsStr), os.ModeAppend) // Imprimindo o arquivo com os grupos criados no algoritmo

	// Error detection in the file operation:
	if err != nil {
		log.Fatalln("The saida.txt file can't be created for some reason...")
	}
}
