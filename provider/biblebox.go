package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	URL       = "https://biblebox-data.turbobytes.net/nvi"
	URL_BOOKS = "https://biblebox-data.turbobytes.net/books_index.json"
)

// Provider para consumo da API do http://biblebox.com/
type Biblebox struct {
}

func (b Biblebox) getDate(book string, chapter string) []interface{} {
	url := fmt.Sprintf("%v/%v/%v.json", URL, getBook(book), chapter)

	res, _ := http.Get(url)

	var date interface{}

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &date)

	m := date.(map[string]interface{})

	return m["verses"].([]interface{})
}

// Obtem o capitulo para o livro solicitado
// a partir da api do biblebox ex.: https://biblebox-data.turbobytes.net/nvi/1/1.json
func (b Biblebox) GetChapter(book string, chapter string) string {
	var out bytes.Buffer
	list := b.getDate(book, chapter)

	for _, v := range list {
		verseNumber := fmt.Sprintf("%v", v.(map[string]interface{})["number"])
		verseText := v.(map[string]interface{})["rawText"].(string)
		out.WriteString("\n" + verseNumber + " " + verseText)
	}

	return out.String()
}

// Obtem o versiculo para o livro e capitulo solicitado
func (b Biblebox) GetVerses(book string, chapter string, verse string) string {
	list := b.getDate(book, chapter)
	verseNumber, _ := strconv.Atoi(verse)
	v := list[verseNumber-1].(map[string]interface{})
	return v["rawText"].(string)
}

func getBook(book string) float64 {
	res, _ := http.Get("https://biblebox-data.turbobytes.net/books_index.json")

	body, _ := ioutil.ReadAll(res.Body)

	var date interface{}

	json.Unmarshal(body, &date)

	font := date.(map[string]interface{})
	return font[strings.ToLower(book)].(float64)
}
