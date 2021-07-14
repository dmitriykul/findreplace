package infrastructure

import (
	"findreplace/pkg/findreplace/app"
	"os"
)

func Deliver(args []string) error {
	var findReplacer app.FindReplacer = &Impl{}
	switch args[0] {
	case "find":
		if len(args) == 2 {
			return findReplacer.FindSubstrInConsoleInput(args[1])
		}
		if len(args) == 3 {
			isDir, err := IsDirectory(args[2])
			if err != nil {
				return err
			}
			if isDir {
				return findReplacer.FindSubstrInDirectory(args[1], args[2])
			}
			return findReplacer.FindSubstrInFile(args[1], args[2])
		}
	case "replace":
		if len(args) == 3 {
			if err := findReplacer.ReplaceSubstrInConsoleInput(args[1], args[2]); err != nil {
				return err
			}
		} else if len(args) == 4 {
			isDir, err := IsDirectory(args[3])
			if err == nil && isDir {
				if err := findReplacer.ReplaceSubstrInDirectory(args[1], args[2], args[3]); err != nil {
					return nil
				}
			} else {
				if err := findReplacer.ReplaceSubstrInFile(args[1], args[2], args[3]); err != nil {
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
