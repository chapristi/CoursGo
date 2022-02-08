package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	tailleDamier   = 9
	symboleJoueur1 = "X"
	symboleJoueur2 = "O"
)

var (
	tableauMorpion = [tailleDamier]string{
		"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9"}
	player1 bool = true
)

func Play() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		var index string

		if player1 == true {
			index = symboleJoueur1
		} else {
			index = symboleJoueur2
		}
		scanner.Scan()
		var text string
		text = scanner.Text()
		if strconv.Atoi(text)-1 > 9 || strconv.Atoi(text)-1 < 1 {
			tableauMorpion[strconv.Atoi(text)-1] = index
			if tableauMorpion[strconv.Atoi(text)-1] != "X" || "O" {
				tableauMorpion[strconv.Atoi(text)-1] = index
				if player1 == true {
					player1 = false
				} else {
					player1 = true
				}
			}
		}
		continue

	}

}

func main() {
	fmt.Printf("test : ", tableauMorpion[0])
	fmt.Println("Hello World")
}
