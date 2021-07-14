package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main(){
	findReplace(os.Args)
}

func findReplace(args []string) {
	args = os.Args[1:]

	// find str file
	// find str dir
	// find str
	// replace str newStr file
	// replace str newStr dir
	// replace str newStr
	switch args[0] {
	case "find":
			if len(args) == 2 {
				findSubstrInConsoleInput(args[1])
			} else if len(args) == 3 {
				if isDirectory(args[2]) {
					findSubstrInDirectory(args[1], args[2])
				} else {
					findSubstrInFile(args[1], args[2])
				}
			}
	case "replace":
			if len(args) == 3 {
				replaceSubstrInConsoleInput(args[1], args[2])
			} else if len(args) == 4 {
				if isDirectory(args[3]) {
					replaceSubstrInDirectory(args[1], args[2], args[3])
				} else {
					replaceSubstrInFile(args[1], args[2], args[3])
				}
			}
	}
}

func isDirectory(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	f.Close()
	if err != nil {
		panic(err)
	}
	if fi.IsDir() {
		return true
	} else {
		return false
	}
}

func findPosition(subStr, text string, count int) int {
	lenC := len(subStr)
	lenS := len(text)
	j := 0

	for i := 0; i <= lenS - lenC; i++ {
		for j = 0; j < lenC && text[i + j] == subStr[j]; j++ {}

		if j == lenC {
			if count- 1 != 0 {
				count--
			} else {
				return i
			}
		}
	}

	return -1
}

// Find a substring in text through Stdin console input
func findSubstrInConsoleInput(str string) {
	var text string
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	text = myScanner.Text()

	/*n := 0
	for i := 1; n != -1; i++ {
		n = findPosition(str, text, i)
		if n != -1 {
			fmt.Println(n)
		}
	}*/

	if findPosition(str, text, 1) != -1 {
		fmt.Printf("%d - %s\n", 1, myScanner.Text())
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
		if findPosition(str, fileScanner.Text(), 1) != -1 {
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
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir()==false {
				findSubstrInFile(str, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

// Replace a substring in text through Stdin console input
func replaceSubstrInConsoleInput(old, new string) {
	var text string
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	text = myScanner.Text()
	fmt.Println(strings.ReplaceAll(text, old, new))
}

func replaceSubstrInFile(str, repStr, file string) {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, str, repStr)
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func replaceSubstrInDirectory(str, repStr, dir string) {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir()==false {
				replaceSubstrInFile(str, repStr, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
