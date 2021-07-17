package infrastructure

import (
	"findreplace/pkg/findreplace/app"
	"io/ioutil"
)

type fileTextStore struct {

}

func NewFileTextStore() app.TextStore {
	return &fileTextStore{}
}

func(t *fileTextStore) StoreText(text []byte, file string) error {
	err := ioutil.WriteFile(file, text, 0644)
	if err != nil {
		return err
	}

	return nil
}
