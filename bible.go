package main

import (
	"bufio"
	"fmt"
	"github.com/weslleyandrade/bibleConsole/provider"
	"os"
)

func main() {
	prov := &provider.Biblebox{}

	lenArgs := len(os.Args)

	if lenArgs > 2 {
		book := os.Args[1]
		chapter := os.Args[2]
		if lenArgs > 3 {
			verse := os.Args[3]
			fmt.Println(prov.GetVerses(book, chapter, verse))

		} else {
			fmt.Println(prov.GetChapter(book, chapter))
		}

	} else {
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
}
