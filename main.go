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
	"time"

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

func barChartPosts(m [][]string) {
	january := 0
	february := 0
	march := 0
	april := 0
	may := 0
	june := 0
	july := 0
	august := 0
	september := 0
	october := 0
	november := 0
	december := 0

	for i := 0; i < len(m); i++ {
		date, err := time.Parse("2006-01-02 15:04:05", m[i][2])
		if err != nil {
			log.Println("Error parsing date:", err)
			continue
		}
		switch date.Month() {
		case time.January:
			january++
		case time.February:
			february++
		case time.March:
			march++
		case time.April:
			april++
		case time.May:
			may++
		case time.June:
			june++
		case time.July:
			july++
		case time.August:
			august++
		case time.September:
			september++
		case time.October:
			october++
		case time.November:
			november++
		case time.December:
			december++
		}
	}
	graph := chart.BarChart{
		Title: "Posts per month",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars: []chart.Value{
			{Value: float64(january), Label: "January"},
			{Value: float64(february), Label: "February"},
			{Value: float64(march), Label: "March"},
			{Value: float64(april), Label: "April"},
			{Value: float64(may), Label: "May"},
			{Value: float64(june), Label: "June"},
			{Value: float64(july), Label: "July"},
			{Value: float64(august), Label: "August"},
			{Value: float64(september), Label: "September"},
			{Value: float64(october), Label: "October"},
			{Value: float64(november), Label: "November"},
			{Value: float64(december), Label: "December"},
		},
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
	}

	f, err := os.Create("barchart.png")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()
	if err := graph.Render(chart.PNG, f); err != nil {
		log.Fatalf("Failed to render graph: %v", err)
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

	f.Println("Total number of words: ", countWords(words))
	f.Println("Total number of unique words: ", countUniqueWords(words))
	displayCharSlice(characters)
	barChartPosts(slice)
}
