package app

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FindParams struct {
	Substr string
	Path string // optional, uses Stdin if empty
}

type ReplaceParams struct {
	Substr string
	Replacement string
	Path string // optional, uses Stdin/Stdout if empty
}

type FindReplacer struct {

}

func (f *FindReplacer) FindSubstr(params FindParams, scanner LineScanner) error {
	if params.Path == "" {
		return f.findSubstrInConsoleInput(params.Substr, scanner)
	}
	isDir, err := isDirectory(params.Path)
	if err != nil {
		return err
	}
	if isDir {
		return f.findSubstrInDirectory(params.Substr, params.Path)
	} else {
		return f.findSubstrInFile(params.Substr, params.Path)
	}


}

func (f *FindReplacer) ReplaceSubstr(params ReplaceParams) error {
	if params.Path == "" {
		return f.replaceSubstrInConsoleInput(params.Substr, params.Replacement)
	}
	isDir, err := isDirectory(params.Path)
	if err != nil {
		return err
	}
	if isDir {
		return f.replaceSubstrInDirectory(params.Substr, params.Replacement, params.Path)
	} else {
		return f.replaceSubstrInDirectory(params.Substr, params.Replacement, params.Path)
	}
}

func isDirectory(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

func (f *FindReplacer) findSubstrInConsoleInput(str string, scanner LineScanner) error {
	lineNo := 0
	for {
		lineNo += 1
		text, err := scanner.ReadConsoleLine(os.Stdin)
		if err != nil {
			return err
		}
		if strings.Contains(text, str) {
			fmt.Printf("%d - %s\n", lineNo, text)
		}
	}

	return nil
}

func (f *FindReplacer) findSubstrInFile(str, path string) error {
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
		if strings.Contains(fileScanner.Text(), str) {
			fmt.Printf("%s:%d - %s\n", separatedPath[len(separatedPath)-1], lineNo, fileScanner.Text())
		}
	}

	if err := fileScanner.Err(); err != nil {
		return err
	}

	return nil
}

func (f *FindReplacer) findSubstrInDirectory(str, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if err := f.findSubstrInFile(str, path); err != nil {
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

func (f *FindReplacer) replaceSubstrInConsoleInput(old, new string) error {
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

func (f *FindReplacer) replaceSubstrInFile(str, repStr, file string) error {
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

func (f *FindReplacer) replaceSubstrInDirectory(str, repStr, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if err := f.replaceSubstrInFile(str, repStr, path); err != nil {
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
