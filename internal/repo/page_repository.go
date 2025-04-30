package repo

import (
	"gowiki/internal/db"
	"gowiki/internal/logger"
	model "gowiki/internal/model"
)

func SavePage(p *model.Page) error {
	_, err := db.DB.Exec(`INSERT INTO pages (title, body) 
		VALUES ($1, $2) 
		ON CONFLICT (title) 
		DO UPDATE SET body = EXCLUDED.body`, p.Title, p.Body)
	if err != nil {
		logger.Error.Printf("Ошибка при сохранении страницы (title=%s: %v): ", p.Title, err)
		return err
	}
	return err
}

func LoadPage(title string) (*model.Page, error) {
	row := db.DB.QueryRow(`SELECT title, body FROM pages WHERE title = $1`, title)
	p := &model.Page{}
	err := row.Scan(&p.Title, &p.Body)
	if err != nil {
		logger.Error.Printf("Ошибка при загрузке страницы (title=%s: %v): ", title, err)
		return nil, err
	}
	return p, nil
}
