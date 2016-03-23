package main

import (
	"bufio"
<<<<<<< HEAD
	"fmt"
	"provider"
=======
>>>>>>> refs/remotes/origin/master
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/weslleyandrade/bibleConsole/provider"
)

func main() {
	prov := &provider.Biblebox{}

	terminal := flag.Bool("t", false, "Modo terminal interativo")
	book := flag.String("l", "", "Livro ex. genesis")
	chapter := flag.String("c", "", "Capitulo ex. 1")
	verse := flag.String("v", "", "Versiculo ex. 2-5")
	query := flag.String("q", "", "Query para consulta ex. -q romanos 1:1-2")
	flag.Parse()

	if *terminal {
		interactive(prov)
	} else {
		if *query != "" {
			book, chapter, verse = queryParse(*query)
		}

		if *verse == "" {
			fmt.Println(prov.GetChapter(*book, *chapter))
		} else {
			fmt.Println(prov.GetVerses(*book, *chapter, *verse))
		}
	}
}

func interactive(prov *provider.Biblebox) {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Livro: ")
		scanner.Scan()
		book := scanner.Text()

		fmt.Printf("%s > Capitulo: ", book)
		scanner.Scan()
		chapter := scanner.Text()

		fmt.Printf("%s > %s > Versiculo: ", book, chapter)
		scanner.Scan()
		verse := scanner.Text()

		if verse == "" {
			fmt.Printf("%s > %s : \n", book, chapter)
			fmt.Println(prov.GetChapter(book, chapter))
		} else {
			fmt.Printf("%s > %s > %s : \n", book, chapter, verse)
			fmt.Println(prov.GetVerses(book, chapter, verse))
		}
	}
}

func queryParse(query string) (*string, *string, *string) {
	var book, chapter, verse *string
	b := strings.Split(query, " ")
	book = &b[0]

	if strings.Contains(b[1], ":") {
		cv := strings.Split(b[1], ":")
		chapter = &cv[0]
		verse = &cv[1]
	} else {
		chapter = &b[1]
		verse = new(string)
	}

	return book, chapter, verse
}
