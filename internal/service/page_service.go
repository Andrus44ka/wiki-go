package service

import (
	"gowiki/internal/logger"
	model "gowiki/internal/model"
	repo "gowiki/internal/repo"
)

type PageService struct{}

// NewPageService конструктор (если захотим DI или тесты)
func NewPageService() *PageService {
	return &PageService{}
}

func (ps *PageService) LoadPage(title string) (*model.Page, error) {
	logger.Info.Printf("Сервис: загрузка страницы %s", title)
	return repo.LoadPage(title)
}

func (ps *PageService) SavePage(page *model.Page) error {
	logger.Info.Printf("Сервис: сохранение страницы %s", page.Title)
	return repo.SavePage(page)
}
