package app

type FindReplacer interface {
	FindPosition(subStr, text string, count int) int
	FindSubstrInConsoleInput(str string) error
	FindSubstrInFile(str, path string) error
	FindSubstrInDirectory(str, dir string) error
	ReplaceSubstrInConsoleInput(old, new string) error
	ReplaceSubstrInFile(str, repStr, file string) error
	ReplaceSubstrInDirectory(str, repStr, dir string) error
}
