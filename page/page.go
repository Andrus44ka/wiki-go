package page

import "os"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
