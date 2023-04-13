package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputIdea = bufio.NewScanner(os.Stdin)
var listaDeIdeas []string

func SaveIdea(lista []string) {
	for {
		fmt.Printf("En que estas pensando? ")
		inputIdea.Scan()
		idea := inputIdea.Text()

		lista = append(lista, idea)

		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		for _, idea := range lista {
			fmt.Printf("- %s \n", idea)
		}
		fmt.Printf("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=\n\n")
	}
}

func main() {
	fmt.Printf("ETP \nEscribe tus pensamientos \nUna app desarrollada por @jheremois en Go.\n.\n.\n.\n\n")
	SaveIdea(listaDeIdeas)

}
