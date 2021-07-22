package app

import (
	"findreplace/pkg/findreplace/infrastructure/workerpool"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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

func (f *FindReplacer) FindSubstr(params FindParams, scanner LineScannerFactory, reporter Reporter) error {
	if params.Path == "" {
		return f.findSubstrInConsoleInput(params.Substr, scanner, reporter)
	}
	isDir, err := isDirectory(params.Path)
	if err != nil {
		return err
	}
	if isDir {
		return f.findSubstrInDirectory(params.Substr, params.Path, reporter, scanner)
	} else {
		return f.findSubstrInFile(params.Substr, params.Path, reporter, scanner)
	}


}

func (f *FindReplacer) ReplaceSubstr(params ReplaceParams, store TextStore, reporter Reporter, scanner LineScannerFactory) error {
	if params.Path == "" {
		return f.replaceSubstrInConsoleInput(params.Substr, params.Replacement, reporter, scanner)
	}
	isDir, err := isDirectory(params.Path)
	if err != nil {
		return err
	}
	if isDir {
		return f.replaceSubstrInDirectory(params.Substr, params.Replacement, params.Path, store)
	} else {
		return f.replaceSubstrInDirectory(params.Substr, params.Replacement, params.Path, store)
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

func (f *FindReplacer) findSubstrInConsoleInput(str string, scanner LineScannerFactory, reporter Reporter) error {
	lineNo := 0
	consoleScanner, _ := scanner.CreateScanner("")
	for {
		lineNo += 1
		exit, text, err := consoleScanner.ReadLine()
		if err != nil {
			return err
		}
		if strings.Contains(text, str) {
			s := strconv.Itoa(lineNo) + " - " + text
			reporter.PrintLine(s)
		}
		if !exit {
			break
		}
	}

	return nil
}

func (f *FindReplacer) findSubstrInFile(str, path string, reporter Reporter, scanner LineScannerFactory) error {
	lineNo := 0
	text := ""
	res := true
	fileScanner, err := scanner.CreateScanner(path)
	if err != nil {
		return err
	}
	for res {
		var err error
		res, text, err = fileScanner.ReadLine()
		if err != nil {
			return err
		}
		lineNo += 1
		if strings.Contains(text, str) {
			text := fileScanner.GetFileName() + ":" + strconv.Itoa(lineNo) + " - " + text
			reporter.PrintLine(text)
		}
	}

	return nil
}

func (f *FindReplacer) findSubstrInDirectory(str, dir string, reporter Reporter, scanner LineScannerFactory) error {
	var allTask []*workerpool.Task
	i := 1
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				task := workerpool.NewTask(func(data interface{}) error {
					return f.findSubstrInFile(str, path, reporter, scanner)
				}, i)
				i += 1
				allTask = append(allTask, task)
			}
			return nil
		})

	if err != nil {
		return err
	}

	pool := workerpool.NewPool(allTask, 4)
	pool.Run()

	return nil
}

func (f *FindReplacer) replaceSubstrInConsoleInput(old, new string, reporter Reporter, scanner LineScannerFactory) error {
	consoleScanner, _ := scanner.CreateScanner("")
	_, text, err := consoleScanner.ReadLine()
	if err != nil {
		return err
	}

	reporter.PrintLine(strings.ReplaceAll(text, old, new))

	return nil
}

func (f *FindReplacer) replaceSubstrInFile(str, repStr, file string, store TextStore) error {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	input = []byte(strings.ReplaceAll(string(input), str, repStr))

	return store.StoreText(input, file)
}

func (f *FindReplacer) replaceSubstrInDirectory(str, repStr, dir string, store TextStore) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == false {
				if err := f.replaceSubstrInFile(str, repStr, path, store); err != nil {
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
