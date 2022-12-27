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

//ROCK = 1 A/X -lose
//Paper = 2 B/Y - draw
//Scisor = 3 C/Z - win
//Win = 6
//Draw = 3

// variant first
func sumScore(data []string) int32 {
	result := int32(0)

	for i := 0; i < len(data); i++ {
		score := int32(0)
		switch data[i] {
		case "A X":
			score = 4
		case "A Y":
			score = 8
		case "A Z":
			score = 3
		case "B X":
			score = 1
		case "B Y":
			score = 5
		case "B Z":
			score = 9
		case "C X":
			score = 7
		case "C Y":
			score = 2
		case "C Z":
			score = 6
		}
		result += score
	}

	return result
}

//ROCK = 1 A/X -lose
//Paper = 2 B/Y - draw
//Scisor = 3 C/Z - win
//Win = 6
//Draw = 3

// variant2
func sumScore2(data []string) int32 {
	result := int32(0)

	for i := 0; i < len(data); i++ {
		score := int32(0)
		switch data[i] {
		case "A X":
			score = 3
		case "A Y":
			score = 4
		case "A Z":
			score = 8
		case "B X":
			score = 1
		case "B Y":
			score = 5
		case "B Z":
			score = 9
		case "C X":
			score = 2
		case "C Y":
			score = 6
		case "C Z":
			score = 7
		}
		result += score
	}

	return result
}

func main() {
	scanner, file := readFile("data.txt")
	elfData := parseElfData(scanner)
	file.Close()
	result := sumScore(elfData)
	result1 := sumScore2(elfData)

	fmt.Println(result)
	fmt.Println(result1)
}
