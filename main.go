package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	//sera appelé par findreplace a chaque ligne lu
	//ligne à traiter
	//old ancien mot
	// new nouveau mot
	//found vrai ou faux si au moins une occurence de trouvé
	//res resultat de la ligne
	//occ nombre d'occurence de old
	isWord := strings.Contains(line, old) // repere si y a un mot dans la ligne
	if isWord {
		nbr := strings.Count(line, old)                  //contient le nombre d'occurence
		res := strings.Replace(line, "go", "python", -1) //remplace
		return isWord, res, nbr
	}
	return false, line, 0

}
func FindReplaceFile(src string, old, new string) (occ int, lines []int, err error) {
	//src le lien du fichier,
	//old l'ancien mot
	//new le nouveau mot
	//occ nombre d'occurence
	//lines les lignes ou old a etait trouvé
	//err erreur de la fonction

	file, err := os.Open(src)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		count += 1
		t := scanner.Text() //contient une ligne
		isFound, res, nbr := ProcessLine(t, old, new)
		if isFound {
			lines[count-1] = count
			occ += nbr
			t = res
		}

	}
	return occ, lines, err

}

func main() {
	occ, lines, err := FindReplaceFile("i.txt", "go", "python")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fmt.Println("les lignes", lines)
	fmt.Println("nombre d'occurence", occ)

}
