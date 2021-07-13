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
	args := os.Args[1:]

	// find str file
	// find str dir
	// find str
	// replace str newStr file
	// replace str newStr dir
	// replace str newStr
	switch args[0] {
		case "find":
			{
			if len(args) == 2 {
				findSubstr(args[1])
			} else if len(args) == 3 {
				if checkPath(args[2]) {
					findSubstrInDirectory(args[1], args[2])
				} else {
					findSubstrInFile(args[1], args[2])
				}
			}
			break
			}
		case "replace":
			{
			if len(args) == 3 {
				replaceSubstr(args[1], args[2])
			} else if len(args) == 4 {
				if checkPath(args[3]) {
					replaceSubstrInDirectory(args[1], args[2], args[3])
				} else {
					replaceSubstrInFile(args[1], args[2], args[3])
				}
			}
			break
			}
	}
}

func checkPath(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	if fi.IsDir() {
		f.Close()
		return true
	} else {
		f.Close()
		return false
	}
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

// Replace a substring in text though Stdin console input
func replaceSubstr(old, new string) {
	var text string
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	text = myScanner.Text()
	fmt.Println(replace(text, old, new))
}

// Replace a substring in file
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

// Replace a substring in directory
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
