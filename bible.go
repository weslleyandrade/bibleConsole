package main

import (
	"fmt"
	"github.com/weslleyandrade/bibleConsole/provider"
)

func main() {
	prov := &provider.Biblebox{}
	for {
		var book string
		var chapter string
		var verse string
		fmt.Print("Livro: ")
		fmt.Scanln(&book)
		fmt.Printf("%s > Capitulo: ", book)
		fmt.Scanln(&chapter)
		fmt.Printf("%s > %s > Versiculo: ", book, chapter)
		fmt.Scanln(&verse)

		if verse == "" {
			fmt.Printf("%s > %s : \n", book, chapter)
			fmt.Println(prov.GetChapter(book, chapter))
		} else {
			fmt.Printf("%s > %s > %s : \n", book, chapter, verse)
			fmt.Println(prov.GetVerses(book, chapter, verse))
		}
	}

}
