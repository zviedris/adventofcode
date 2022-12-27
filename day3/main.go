package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(name string) (*bufio.Reader, *os.File) {
	//open and read the input file
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()

	scanner := bufio.NewReader(file)
	return scanner, file
}

func parseElfData(scanner *bufio.Reader) []string {
	//parse lines
	var elfData []string
	var err error = nil
	var str []byte
	for err == nil {
		str, _, err = scanner.ReadLine()
		//if new line is read
		if err == nil {
			//translate to string
			str1 := (strings.TrimRight(string(str), "\r\n"))
			elfData = append(elfData, str1)
		}

	}
	return elfData
}

func letterToValue(letter byte) int {
	if letter < 92 {
		return int(letter) - 64 + 26
	} else {
		return int(letter) - 96
	}
}

func calculateResultInd(data string) int {
	part1 := data[:len(data)/2]
	part2 := string(data[len(data)/2:])

	for i := 0; i < len(part1); i++ {
		if strings.Contains(part2, string(part1[i])) {
			return letterToValue(part1[i])
		}
	}

	return 0
}

func calculateResult(data []string) int {
	result := 0

	for i := 0; i < len(data); i++ {
		result += calculateResultInd(data[i])
	}

	return result
}

func main() {
	scanner, file := readFile("data.txt")
	elfData := parseElfData(scanner)
	file.Close()
	result := calculateResult(elfData)

	fmt.Println(result)
	//fmt.Println(result1)
}
