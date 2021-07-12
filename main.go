package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args[1:]
	strToFind := args[1]
	source := args[2]
	n := 0
	fmt.Println(args)
	for i := 1; n != -1; i++ {
		n = pos(strToFind, source, i)
		fmt.Println(n)
	}
	//if args[0] == "find" {
	//	stringToFind = args[1]
	//}
	//else if args[0] == "replace" {
	//	stringToFind = args[1]
	//	stringToReplace = args[2]
	//}
	//fmt.Printf("%B", findSubstr("hello"))
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

// Find a substring in text through Stdin console input
func findSubstr(str string) {
	var text string
	fmt.Fscan(os.Stdin, &text)


}

// Find a substring in file and print through Stdout
func findSubstrInFile(str, file string) {

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
