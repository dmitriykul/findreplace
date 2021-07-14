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
	if err := findReplace(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func findReplace(args []string) error {
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
				if err := findSubstrInConsoleInput(args[1]); err != nil {
					return err
				}
			} else if len(args) == 3 {
				isDir, err := isDirectory(args[2])
				if err == nil && isDir {
					if err := findSubstrInDirectory(args[1], args[2]); err != nil {
						return err
					}
				} else {
					if err := findSubstrInFile(args[1], args[2]); err != nil {
						return err
					}
				}
			}
	case "replace":
			if len(args) == 3 {
				if err := replaceSubstrInConsoleInput(args[1], args[2]); err != nil {
					return err
				}
			} else if len(args) == 4 {
				isDir, err := isDirectory(args[3])
				if err == nil && isDir {
					if err := replaceSubstrInDirectory(args[1], args[2], args[3]); err != nil {
						return nil
					}
				} else {
					if err := replaceSubstrInFile(args[1], args[2], args[3]); err != nil {
						return err
					}
				}
			}
	}
	return nil
}

func isDirectory(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	fi, err := f.Stat()
	f.Close()
	if err != nil {
		return false, err
	}
	if fi.IsDir() {
		return true, nil
	} else {
		return false, nil
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
func findSubstrInConsoleInput(str string) error {
	var text string
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	text = myScanner.Text()

	if err := myScanner.Err(); err != nil {
		return err
	}

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

	return nil
}

// Find a substring in file and print through Stdout
func findSubstrInFile(str, path string) error {
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
		return err
	}
	file.Close()

	return nil
}

// Find a substring in directory and print through Stdout
func findSubstrInDirectory(str, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir()==false {
				if err := findSubstrInFile(str, path); err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}

	return nil
}

// Replace a substring in text through Stdin console input
func replaceSubstrInConsoleInput(old, new string) error {
	var text string
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	text = myScanner.Text()

	if err := myScanner.Err(); err != nil {
		return err
	}
	fmt.Println(strings.ReplaceAll(text, old, new))

	return nil
}

func replaceSubstrInFile(str, repStr, file string) error {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, str, repStr)
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

func replaceSubstrInDirectory(str, repStr, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir()==false {
				if err := replaceSubstrInFile(str, repStr, path); err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}

	return nil
}
