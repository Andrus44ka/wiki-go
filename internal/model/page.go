// только структура Page (бизнес-сущность)

package wiki

type Page struct {
	ID    int
	Title string
	Body  []byte
}
