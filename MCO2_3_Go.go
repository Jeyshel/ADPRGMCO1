// ********************
// Last names: Gaffud
// Language: Go
// Paradigm(s): Imperative
// ********************

package main

import (
	"encoding/csv"
	f "fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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

func isStopWord(word string) string {
	stopWords := []string{"a", "an", "and", "are", "as", "at", "be", "by", "for", "from", "has", "he", "in", "is", "it", "its", "of", "on", "that", "the", "to", "was", "were", "will", "with"}
	for i := 0; i < len(stopWords); i++ {
		if word == stopWords[i] {
			return "Y"
		}
	}
	return ""
}

func displayCharSlice(m []rKeyVal) {
	f.Println("Character frequency count: ")
	for i := 0; i < len(m); i++ {
		f.Println(i+1, string(m[i].Key), m[i].Value)
	}
}

func displayMapSlice(m []sKeyVal) {
	for i := 0; i < len(m); i++ {
		f.Println(i+1, ":", isStopWord(m[i].Key), ":", m[i].Key, ":", m[i].Value)

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
		date, _ := time.Parse("2006-01-02 15:04:05", m[i][2])

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

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Posts per month",
	}))

	bar.SetXAxis([]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}).
		AddSeries("Posts", []opts.BarData{
			{Value: january},
			{Value: february},
			{Value: march},
			{Value: april},
			{Value: may},
			{Value: june},
			{Value: july},
			{Value: august},
			{Value: september},
			{Value: october},
			{Value: november},
			{Value: december},
		})
	f, _ := os.Create("bar.html")
	bar.Render(f)

}

func wordCloudHTML(words map[string]int) {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Word Cloud",
	}))

	items := make([]opts.WordCloudData, 0, len(words))
	for word, count := range words {
		items = append(items, opts.WordCloudData{Name: word, Value: count})
	}

	wc.AddSeries("wordcloud", items)
	f, _ := os.Create("wordcloud.html")
	wc.Render(f)
}

func topNWords(wordMap map[string]int, n int) map[string]int {
	sortedWords := sortStringMapByValueDesc(wordMap)
	topWords := make(map[string]int)
	for i := 0; i < n && i < len(sortedWords); i++ {
		topWords[sortedWords[i].Key] = sortedWords[i].Value
	}
	return topWords
}

func isSymbol(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsDigit(c)
}

func symbolPieChart(symbols map[rune]int) {
	pie := charts.NewPie()

	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Symbol frequency",
	}))

	items := make([]opts.PieData, 0, len(symbols))
	for symbol, count := range symbols {
		if isSymbol(symbol) {
			items = append(items, opts.PieData{Name: string(symbol), Value: count})
		}
	}

	pie.AddSeries("Symbols", items)
	f, _ := os.Create("pie.html")
	pie.Render(f)
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
	displayMapSlice(sortStringMapByValueDesc(mapWords(words)))
	displayCharSlice(characters)

	barChartPosts(slice)
	symbolPieChart(convertToChar(lines))

	wordMap := mapWords(words)
	topWords := topNWords(wordMap, 20)
	wordCloudHTML(topWords)
}
