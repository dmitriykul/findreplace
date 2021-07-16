package infrastructure

import "findreplace/pkg/findreplace/app"

type textStore struct {

}

func NewTextStore() app.TextStore {
	return &textStore{}
}

func(t *textStore) StoreText(text, file string) error {

	return nil
}
