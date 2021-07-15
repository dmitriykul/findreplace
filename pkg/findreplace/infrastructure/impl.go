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

}

func (i *Impl) FindSubstr(params app.FindParams) error {
	if params.Path == "" {
		return i.findSubstrInConsoleInput(params.Substr)
	}
	isDir, err := isDirectory(params.Path)
	if err != nil {
		return err
	}
	if isDir {
		return i.findSubstrInDirectory(params.Substr, params.Path)
	} else {
		return i.findSubstrInFile(params.Substr, params.Path)
	}


}

func (i *Impl) ReplaceSubstr(params app.ReplaceParams) error {
	if params.Path == "" {
		return i.replaceSubstrInConsoleInput(params.Substr, params.Replacement)
	}
	isDir, err := isDirectory(params.Path)
	if err != nil {
		return err
	}
	if isDir {
		return i.replaceSubstrInDirectory(params.Substr, params.Replacement, params.Path)
	} else {
		return i.replaceSubstrInDirectory(params.Substr, params.Replacement, params.Path)
	}
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
	return fi.IsDir(), nil
}

func (i *Impl) findPosition(subStr, text string, count int) int {
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

func (i *Impl) findSubstrInConsoleInput(str string) error {
	scanner := bufio.NewScanner(os.Stdin)
	lineNo := 0
	for scanner.Scan() {
		lineNo += 1
		text := scanner.Text()
		if i.findPosition(str, text, 1) != -1 {
			fmt.Printf("%d - %s\n", lineNo, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (i *Impl) findSubstrInFile(str, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	lineNo := 0
	separatedPath := strings.Split(path, "\\")
	for fileScanner.Scan() {
		lineNo += 1
		if i.findPosition(str, fileScanner.Text(), 1) != -1 {
			fmt.Printf("%s:%d - %s\n", separatedPath[len(separatedPath)-1], lineNo, fileScanner.Text())
		}
	}

	if err := fileScanner.Err(); err != nil {
		return err
	}

	return nil
}

func (i *Impl) findSubstrInDirectory(str, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if err := i.findSubstrInFile(str, path); err != nil {
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

func (i *Impl) replaceSubstrInConsoleInput(old, new string) error {
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

func (i *Impl) replaceSubstrInFile(str, repStr, file string) error {
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

func (i *Impl) replaceSubstrInDirectory(str, repStr, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if err := i.replaceSubstrInFile(str, repStr, path); err != nil {
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