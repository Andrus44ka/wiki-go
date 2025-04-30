// только структура Page (бизнес-сущность)

package model

type Page struct {
	ID    int
	Title string
	Body  []byte
}
