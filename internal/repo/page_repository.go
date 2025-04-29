package wiki

import (
	"gowiki/internal/db"
	model "gowiki/internal/model"
)

func SavePage(p *model.Page) error {
	_, err := db.DB.Exec(`INSERT INTO pages (title, body) 
		VALUES ($1, $2) 
		ON CONFLICT (title) 
		DO UPDATE SET body = EXCLUDED.body`, p.Title, p.Body)
	return err
}

func LoadPage(title string) (*model.Page, error) {
	row := db.DB.QueryRow(`SELECT title, body FROM pages WHERE title = $1`, title)
	p := &model.Page{}
	err := row.Scan(&p.Title, &p.Body)
	if err != nil {
		return nil, err
	}
	return p, nil
}
