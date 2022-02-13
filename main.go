package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine searches for old in line to replace it by new
// It returns found=true, if the pattern was found, res with resulting sting
// and occ with the numbre of occurence of old
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	var oldLower string = strings.ToLower(old)
	var newLower string = strings.ToLower(new)
	res = line
	//strings.Contains trouve si il y a le vieux mot dans la ligne données
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		//on incremente au cas ou le mot commencerais par une majuscule
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		//on incremente au cas ou le mot commencerais par une majuscule
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func FindReplaceFile(src string, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()
	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		println(res)
		lineIdx++
	}
	return occ, lines, nil
}

func main() {
	old := "Go"
	newWord := "Python"
	occ, lines, err := FindReplaceFile("i.txt", old, newWord)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("== summary")
	defer fmt.Println("== End of summary")
	fmt.Printf("number %v\n", occ)
	fmt.Printf("number of lines %d\n", len(lines))
	fmt.Printf("Lines : [ ")

	for _, l := range lines {
		fmt.Printf("%v - ", l)
	}
	fmt.Printf("]\n")
}
