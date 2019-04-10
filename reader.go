/**
 * Linguagens de Programação - Trabalho 1
 *
 * Alan Herculano Diniz
 *
 * Resolver problema de agrupamento com o algoritmo dado em aula
 *
 * reader.go: manipulação de arquivos para entrada e saída
 *
 * ***********************************************************************************
 *
 * Coisas a fazer:
 *
 * 1 - Criar função para transformar a string com coordenadas em uma matriz de números
 */

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"fmt"
)

/**
 * Lendo o arquivo de entrada
 *
 * Esta função lê o arquivo de entrada e cria uma representação primitiva
 * dos dados que precisam ser manipulados
 *
 * Entrada: localização do arquivo de entrada
 *
 * Saída: string com o conteúdo do arquivo
 */
func readFile(loc string) ([]string, int) {

	file, err := ioutil.ReadFile(loc) // Lendo o arquivo de entrada

	// Verificação de erro na leitura do arquivo:
	if err != nil {
		log.Fatalln("The passed file can't be read for some reason...")
	}

	str := string(file) // Convertendo o conteúdo para uma string
	lines := 0 // Número de linhas na string
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			lines++
		}
	}

	result := make([]string, lines)
	for i := 0; i < lines; i++ {
		fmt.Sscanf(str, "%s\n", result[i])
	}

	return result, lines
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

/**
 * Convertendo string de entrada em uma matriz de números
 *
 * Entrada: string com as coordenadas dos pontos
 *
 * Saída: matriz em que cada posição é um vetor com as coordenadas de cada ponto
 */
func parseInputString(input string, lines int) [][]float64 {

	var result [][]float64 = [][]float64{{1, 1}, {1, 1}}

	

	return result

}
