package provider

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (b Biblebox) getDate(book string, chapter string) ([]interface{}, error) {
	var verses []interface{}
	nBook, err := getBook(book)
	if err != nil {
		return verses, err
	}

	url := fmt.Sprintf("%v/%v/%v.json", URL, nBook, chapter)

	res, _ := http.Get(url)

	if res.StatusCode != 200 {
		return verses, errors.New("Capitulo nao encontrado!")
	}

	body, _ := ioutil.ReadAll(res.Body)

	var date interface{}
	json.Unmarshal(body, &date)

	m := date.(map[string]interface{})
	verses = m["verses"].([]interface{})

	return verses, nil
}

// Obtem o capitulo para o livro solicitado
// a partir da api do biblebox ex.: https://biblebox-data.turbobytes.net/nvi/1/1.json
func (b Biblebox) GetChapter(book string, chapter string) string {
	var out bytes.Buffer

	list, err := b.getDate(book, chapter)

	if err != nil {
		return err.Error()
	}

	for _, v := range list {
		verseNumber := fmt.Sprintf("%v", v.(map[string]interface{})["number"])
		verseText := v.(map[string]interface{})["rawText"].(string)
		out.WriteString("\n" + verseNumber + " " + verseText)
	}

	return out.String()
}

// Obtem o versiculo para o livro e capitulo solicitado
func (b Biblebox) GetVerses(book string, chapter string, verse string) string {

	list, err := b.getDate(book, chapter)
	if err != nil {
		return err.Error()
	}
	verses := strings.Split(verse, "-")
	minorVerse, _ := strconv.Atoi(verses[0])
	majorVerse, _ := strconv.Atoi(verses[1])

	verseIndex := majorVerse - 1
	if len(list) < verseIndex || minorVerse > majorVerse {
		return "Versiculo nao encontrado, ou formato inconsistente"
	}

	var out bytes.Buffer

	for i := minorVerse; i <= majorVerse; i++ {
		v := list[i-1]
		verseNumber := fmt.Sprintf("%v", v.(map[string]interface{})["number"])
		out.WriteString("\n" + verseNumber + " " + v.(map[string]interface{})["rawText"].(string))
	}

	return out.String()
}

func getBook(book string) (string, error) {
	var nBook float64
	var err error

	res, _ := http.Get("https://biblebox-data.turbobytes.net/books_index.json")

	body, _ := ioutil.ReadAll(res.Body)

	var date interface{}

	json.Unmarshal(body, &date)

	font := date.(map[string]interface{})

	auxBook := font[strings.ToLower(book)]

	if auxBook == nil {
		err = errors.New("Livro nao encontrado!")
	} else {
		nBook = auxBook.(float64)
	}

	return fmt.Sprint(nBook), err

}
