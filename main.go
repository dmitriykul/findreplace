package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var text string
	args := os.Args[1:]
	strToFind := args[1]
	if len(args) == 2 {
		findSubstr(strToFind)
	} else if len(args) == 3 {
		findSubstrInFile(strToFind, args[2])
	}
	fmt.Println(replace(text, strToFind, "ra"))
}

func pos(c, s string, n int) int {
	var lenC, lenS, j int
	lenC = len(c)
	lenS = len(s)

	for i := 0; i <= lenS - lenC; i++ {
		for j = 0; j < lenC && s[i + j] == c[j]; j++ {}

		if j == lenC {
			if n - 1 != 0 {
				n--
			} else {
				return i
			}
		}
	}

	return -1
}

func replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Find a substring in text through Stdin console input
func findSubstr(str string) {
	var text string
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	text = myScanner.Text()

	n := 0
	for i := 1; n != -1; i++ {
		n = pos(str, text, i)
		if n != -1 {
			fmt.Println(n)
		}
	}

}

// Find a substring in file and print through Stdout
func findSubstrInFile(str, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var i int
	separatedPath := strings.Split(path, "\\")
	for fileScanner.Scan(){
		i+=1
		if pos(str, fileScanner.Text(), 1) != -1 {
			fmt.Printf("%s:%d - %s\n", separatedPath[len(separatedPath)-1], i, fileScanner.Text())
		}
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s", err)
	}
	file.Close()
}

// Find a substring in directory and print through Stdout
func findSubstrInDirectory(str, dir string) {

}

// Replace a substring in text though Stdin console input
func replaceSubstr(str, repStr string) {

}

// Replace a substring in file
func replaceSubstrInFile(str, repStr, file string) {

}

// Replace a substring in directory
func replaceSubstrInDirectory(str, repStr, dir string) {

}
