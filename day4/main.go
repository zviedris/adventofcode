package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func calculateResultInd(a int, b int, c int, d int) int {

	if a <= c && b >= d {
		return 1
	}

	if a >= c && b <= d {
		return 1
	}

	return 0
}

func calculateResultInd1(a int, b int, c int, d int) int {

	//pairs do not overlap
	if a > d || c > b {
		return 0
	}

	return 1
}

func parseLine(line string) (int, int, int, int) {
	a := strings.Split(line, ",")
	c := strings.Split(a[0], "-")
	d := strings.Split(a[1], "-")

	v1, _ := strconv.ParseInt(c[0], 10, 64)
	v2, _ := strconv.ParseInt(c[1], 10, 64)
	v3, _ := strconv.ParseInt(d[0], 10, 64)
	v4, _ := strconv.ParseInt(d[1], 10, 64)

	return int(v1), int(v2), int(v3), int(v4)
}

func calculateResult(data []string) int {
	result := 0

	for i := 0; i < len(data); i++ {
		result += calculateResultInd1(parseLine(data[i]))
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
