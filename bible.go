package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type bibleBook struct {
	Name    string
	Version string
}

type bibleChapter struct {
	Number int
	Verses int
}

type bibleVerse struct {
	Number int
	Text   string
}

type bibleResult struct {
	Book    bibleBook
	Verses  []bibleVerse
	Number  int
	Chapter interface{}
	Text    string
	Msg     string
}

type bibleAPI struct{}

func (*bibleAPI) readReference(chapter string, verse string, version string) bibleResult {
	resp, err := http.Get(
		fmt.Sprintf("https://www.abibliadigital.com.br/api/verses/%s/%s/%s",
			version, chapter, strings.Replace(verse, ":", "/", 1)))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result bibleResult
	json.Unmarshal([]byte(string(body)), &result)
	return result
}

func (*bibleAPI) printResult(result bibleResult) {
	if result.Msg != "" {
		log.Fatal(result.Msg)
		os.Exit(1)
	}

	var chapter int
	switch result.Chapter.(type) {
	case float64:
		chapter = int(result.Chapter.(float64))
	default:
		chapter = int(result.Chapter.(map[string]interface{})["number"].(float64))
	}

	titleColor := color.New(color.FgYellow).Add(color.Underline)
	titleColor.Printf("%s %d (%s)\n", result.Book.Name, chapter, result.Book.Version)

	textColor := color.New(color.FgCyan)
	if result.Text != "" {
		textColor.Printf("%d. \"%s\"\n", result.Number, result.Text)
	} else {
		for _, value := range result.Verses {
			textColor.Printf("%d. \"%s\"\n", value.Number, value.Text)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("bible.go - a command-line interface to aBibliaDigital API")
	}

	if len(os.Args) <= 2 {
		fmt.Println("Usage: bible.go book-code verse-id (e.g. mt 6:23)")
		os.Exit(1)
	}

	var bibleAPI bibleAPI
	var result bibleResult
	if len(os.Args) == 3 {
		result = bibleAPI.readReference(os.Args[1], os.Args[2], "acf")
	} else {
		result = bibleAPI.readReference(os.Args[1], os.Args[2], os.Args[3])
	}
	bibleAPI.printResult(result)
}
