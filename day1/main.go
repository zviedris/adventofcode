package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func parseElfData(scanner *bufio.Reader) [][]int32 {
	//parse lines
	var elfData [][]int32
	var elf []int32 = make([]int32, 0)
	var err error = nil
	var str []byte
	for err == nil {
		str, _, err = scanner.ReadLine()
		//if new line is read
		if err == nil {
			//translate to string
			str1 := (strings.TrimRight(string(str), "\r\n"))
			//if empty then elf data ends and add
			if str1 == "" {
				elfData = append(elfData, elf)
				elf = make([]int32, 0)
			} else {
				//add item to elf data
				dataItem, _ := strconv.ParseInt(str1, 10, 32)
				elf = append(elf, int32(dataItem))
			}
		}

	}
	return elfData
}

func sumElfs(elfData [][]int32) []int32 {
	var result []int32

	for i := 0; i < len(elfData); i++ {
		current := int32(0)
		for j := 0; j < len(elfData[i]); j++ {
			current += elfData[i][j]
		}
		result = append(result, current)
	}

	return result
}

func sumTop(data []int32, top int) int32 {
	result := int32(0)

	for i := 0; i < top; i++ {
		result += data[i]
	}

	return result
}

func main() {
	scanner, file := readFile("data.txt")
	elfData := parseElfData(scanner)
	file.Close()
	elfSums := sumElfs(elfData)

	sort.SliceStable(elfSums, func(i, j int) bool {
		return elfSums[i] > elfSums[j]
	})

	result := sumTop(elfSums, 3)

	fmt.Println(result)
}
