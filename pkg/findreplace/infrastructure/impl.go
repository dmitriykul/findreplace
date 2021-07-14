package infrastructure

import (
	"bufio"
	"findreplace/pkg/findreplace/app"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Impl struct {
	args []string
}

func NewImpl(args []string) *Impl {
	i := new(Impl)
	i.args = args

	return i
}

func (i Impl) FindPosition(subStr, text string, count int) int {
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

func (i Impl) FindSubstrInConsoleInput(str string) error {
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

func (i Impl) FindSubstrInFile(str, path string) error {
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

func (i Impl) FindSubstrInDirectory(str, dir string) error {
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

func (i Impl) ReplaceSubstrInConsoleInput(old, new string) error {
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

func (i Impl) ReplaceSubstrInFile(str, repStr, file string) error {
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

func (i Impl) ReplaceSubstrInDirectory(str, repStr, dir string) error {
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