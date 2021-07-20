package app_test

import (
	"findreplace/pkg/findreplace/app"
	"findreplace/pkg/findreplace/infrastructure"
	"io/ioutil"
	"os"
	"testing"
)

// main find pa
// main find pa file.txt
// main find pa dir

// main replace pa ra
// main replace pa ra file.txt
// main replace pa ra dir

type TestFindParams struct {
	params app.FindParams
	scanner app.LineScanner
	reporter app.Reporter
}

type TestReplaceParams struct {
	params app.ReplaceParams
	store app.TextStore
}

func TestFindInConsoleInput(t *testing.T) {
	
}
func TestReplaceInConsoleInput(t *testing.T) {

}

func TestFindInfile(t *testing.T) {
	var findReplacer app.FindReplacer
	var params TestFindParams
	params.initializeTestFindParams("file")

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	findReplacer.FindSubstr(params.params, params.scanner, params.reporter)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != "file.txt:1 - papapa\n" {
		t.Fatalf("Expected 'file.txt:1 - papapa' but was: %s", out)
	}
}

func TestReplaceInFile(t *testing.T) {
	var findReplacer app.FindReplacer
	var params TestReplaceParams
	params.initializeTestReplaceParams("file")

	findReplacer.ReplaceSubstr(params.params, params.store)
	input, _ := ioutil.ReadFile("test/file.txt")
	if string(input) != "rararara" {
		t.Fatalf("Expected 'rararara' but was: %s", input)
	}
}

func TestFindInDir(t *testing.T) {
	var findReplacer app.FindReplacer
	var params TestFindParams
	params.initializeTestFindParams("dir")

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	findReplacer.FindSubstr(params.params, params.scanner, params.reporter)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != "file.txt:1 - papapa\nfile2.txt:1 - pa\n" {
		t.Fatalf("Expected 'file.txt:1 - papapa\nfile2.txt:1 - pa' but was: %s", out)
	}
}

func TestReplaceInDir(t *testing.T) {
	var findReplacer app.FindReplacer
	var params TestReplaceParams
	params.initializeTestReplaceParams("dir")
	findReplacer.ReplaceSubstr(params.params, params.store)
	input, _ := ioutil.ReadFile("test/file.txt")
	if string(input) != "rarara" {
		t.Fatalf("Expected 'rarara' but was: %s", input)
	}
	input, _ = ioutil.ReadFile("test/subTest/file2.txt")
	if string(input) != "ra" {
		t.Fatalf("Expected 'ra' but was: %s", input)
	}
}

func (t *TestFindParams) initializeTestFindParams(source string) {
	t.params.Substr = "pa"
	if source == "file" {
		os.Mkdir("test", 0777)
		os.Create("test/file.txt")
		os.WriteFile("test/file.txt", []byte("papapa"), 0777)
		t.params.Path = "test/file.txt"
		t.scanner, _ = infrastructure.NewFileScanner("test/file.txt")
		t.reporter = infrastructure.NewReporter()
	} else if source == "dir" {
		os.MkdirAll("test/subTest", 0777)
		os.Create("test/file.txt")
		os.WriteFile("test/file.txt", []byte("papapa"), 0777)
		os.Create("test/subTest/file2.txt")
		os.WriteFile("test/subTest/file2.txt", []byte("pa"), 0777)
		t.params.Path = "test"
		t.scanner, _ = infrastructure.NewFileScanner("test/file.txt")
		t.reporter = infrastructure.NewReporter()
	} else if source == "console" {
		t.scanner = nil
		t.reporter = infrastructure.NewReporter()
	}
}

func (t *TestReplaceParams) initializeTestReplaceParams(source string) {
	t.params.Substr = "pa"
	t.params.Replacement = "ra"
	if source == "file" {
		os.Mkdir("test", 0777)
		os.Create("test/file.txt")
		os.WriteFile("test/file.txt", []byte("papapapa"), 0777)
		t.params.Path = "test/file.txt"
		t.store = infrastructure.NewFileTextStore()
	} else if source == "dir" {
		os.MkdirAll("test/subTest", 0777)
		os.Create("test/file.txt")
		os.WriteFile("test/file.txt", []byte("papapa"), 0777)
		os.Create("test/subTest/file2.txt")
		os.WriteFile("test/subTest/file2.txt", []byte("pa"), 0777)
		t.params.Path = "test"
		t.store = infrastructure.NewFileTextStore()
	}
}