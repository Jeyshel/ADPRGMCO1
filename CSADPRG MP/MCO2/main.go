package main

import (
	"encoding/csv"
	"fmt"
	f "fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"github.com/wcharczuk/go-chart"
)

type sKeyVal struct {
	Key   string
	Value int
}

type rKeyVal struct {
	Key   rune
	Value int
}

func convertToSlice(s [][]string) []string {
	var slice []string
	for i := 0; i < len(s); i++ {
		for j := 3; j < len(s[i])-1; j++ {
			slice = append(slice, s[i][j])
		}
	}
	return slice
}

func convertToWords(str []string) []string {
	var slice []string
	for i := 0; i < len(str); i++ {
		slice = append(slice, strings.Fields(str[i])...)
	}
	return slice
}

func convertToChar(str []string) map[rune]int {
	char := make(map[rune]int)
	for i := 0; i < len(str); i++ {
		for _, c := range str[i] {
			char[c]++
		}
	}
	return char
}

func sortRuneMapByValueDesc(m map[rune]int) []rKeyVal {
	var mapSlice []rKeyVal
	for k, v := range m {
		mapSlice = append(mapSlice, rKeyVal{k, v})
	}

	sort.Slice(mapSlice, func(i, j int) bool {
		return mapSlice[i].Value > mapSlice[j].Value
	})

	return mapSlice
}
func countWords(str []string) int {
	return len(str)
}

func mapWords(str []string) map[string]int {
	myMap := make(map[string]int)
	for i := 0; i < len(str); i++ {
		myMap[str[i]]++
	}
	return myMap
}

func countUniqueWords(str []string) int {
	return len(mapWords(str))
}

func sortStringMapByValueDesc(m map[string]int) []sKeyVal {
	var mapSlice []sKeyVal
	for k, v := range m {
		mapSlice = append(mapSlice, sKeyVal{k, v})
	}

	sort.Slice(mapSlice, func(i, j int) bool {
		return mapSlice[i].Value > mapSlice[j].Value
	})

	return mapSlice
}

func displaySlice(str []string) {
	for i := 0; i < len(str); i++ {
		f.Println(str[i])
	}
}

func displayMap(m map[string]int) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func displayTop20(arr []sKeyVal) {
	f.Println("The top 20 words are: ")
	for i, kv := range arr {
		if i >= 20 {
			break
		}
		f.Printf("Key: %s, Value: %d\n", kv.Key, kv.Value)
	}
}

func displayCharSlice(m []rKeyVal) {
	f.Println("Character frequency count: ")
	for _, kv := range m {

		f.Printf("Key: %c, Value: %d\n", kv.Key, kv.Value)
	}
}

func main() {
	var filename string
	f.Println("Enter file name: ")
	f.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		f.Println("File error")
	}
	defer file.Close()

	var slice [][]string
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		slice = append(slice, record)
	}

	lines := convertToSlice(slice)
	words := convertToWords(lines)
	characters := sortRuneMapByValueDesc(convertToChar(lines))
	wordMap := mapWords(words)
	sortedList := sortStringMapByValueDesc(wordMap)

	f.Println("Total number of words: ", countWords(words))
	f.Println("Total number of unique words: ", countUniqueWords(words))
	displayCharSlice(characters)
	displayTop20(sortedList)

}
