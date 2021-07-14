package app

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func FindReplace(args []string) error {
	// find str file
	// find str dir
	// find str
	// replace str newStr file
	// replace str newStr dir
	// replace str newStr
	switch args[0] {
	case "find":
		if len(args) == 2 {
			return FindSubstrInConsoleInput(args[1])
		}
		if len(args) == 3 {
			isDir, err := IsDirectory(args[2])
			if err != nil {
				return err
			}
			if isDir {
				return FindSubstrInDirectory(args[1], args[2])
			}
			return FindSubstrInFile(args[1], args[2])
		}
	case "replace":
		if len(args) == 3 {
			if err := ReplaceSubstrInConsoleInput(args[1], args[2]); err != nil {
				return err
			}
		} else if len(args) == 4 {
			isDir, err := IsDirectory(args[3])
			if err == nil && isDir {
				if err := ReplaceSubstrInDirectory(args[1], args[2], args[3]); err != nil {
					return nil
				}
			} else {
				if err := ReplaceSubstrInFile(args[1], args[2], args[3]); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func IsDirectory(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	fi, err := f.Stat()
	f.Close()
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

func FindPosition(subStr, text string, count int) int {
	lenC := len(subStr)
	lenS := len(text)
	j := 0

	for i := 0; i <= lenS-lenC; i++ {
		for j = 0; j < lenC && text[i+j] == subStr[j]; j++ {
		}

		if j == lenC {
			if count-1 != 0 {
				count--
			} else {
				return i
			}
		}
	}

	return -1
}

// Find a substring in text through Stdin console input
func FindSubstrInConsoleInput(str string) error {
	scanner := bufio.NewScanner(os.Stdin)
	lineNo := 0
	for scanner.Scan() {
		lineNo += 1
		text := scanner.Text()
		if FindPosition(str, text, 1) != -1 {
			fmt.Printf("%d - %s\n", lineNo, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Find a substring in file and print through Stdout
func FindSubstrInFile(str, path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	lineNo := 0
	separatedPath := strings.Split(path, "\\")
	for fileScanner.Scan() {
		lineNo += 1
		if FindPosition(str, fileScanner.Text(), 1) != -1 {
			fmt.Printf("%s:%d - %s\n", separatedPath[len(separatedPath)-1], lineNo, fileScanner.Text())
		}
	}

	if err := fileScanner.Err(); err != nil {
		return err
	}
	file.Close()

	return nil
}

// Find a substring in directory and print through Stdout
func FindSubstrInDirectory(str, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if err := FindSubstrInFile(str, path); err != nil {
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
func ReplaceSubstrInConsoleInput(old, new string) error {
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

func ReplaceSubstrInFile(str, repStr, file string) error {
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

func ReplaceSubstrInDirectory(str, repStr, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if err := ReplaceSubstrInFile(str, repStr, path); err != nil {
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
