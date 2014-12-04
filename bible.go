package main

import (
	"bufio"
	"fmt"
	"github.com/weslleyandrade/bibleConsole/provider"
	"os"
)

func main() {
	prov := &provider.Biblebox{}
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
