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

func calculateResultInd(data string, data1 string, data2 string) int {

	for i := 0; i < len(data); i++ {
		if strings.Contains(data1, string(data[i])) && strings.Contains(data2, string(data[i])) {
			return letterToValue(data[i])
		}
	}

	return 0
}

func calculateResult(data []string) int {
	result := 0

	for i := 0; i < len(data); i += 3 {
		result += calculateResultInd(data[i], data[i+1], data[i+2])
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
