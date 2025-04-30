package wiki

import (
	"gowiki/internal/logger"
	model "gowiki/internal/model"
	service "gowiki/internal/service"
	"html/template"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var pageService = service.NewPageService()

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

var templates = template.Must(template.ParseFiles("web/temp/edit.html", "web/temp/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, page *model.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", page)
	if err != nil {
		logger.Error.Printf("Ошибка рендера шаблона %s: %v", tmpl, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	logger.Info.Printf("Получен запрос на просмотр страницы: %v", title)
	page, err := pageService.LoadPage(title)
	if err != nil {
		logger.Warn.Printf("Страница %s не найдена, перенаправление на /edit", title)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	logger.Info.Printf("Страница %s успешно загружена", title)
	renderTemplate(w, "view", page)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	logger.Info.Printf("Получен запрос на редактирование страницы: %s", title)

	page, err := pageService.LoadPage(title)
	if err != nil {
		logger.Warn.Printf("Страница %s не найдена, создается новая", title)
		page = &model.Page{Title: title}
	}
	renderTemplate(w, "edit", page)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	logger.Info.Printf("Получен запрос на сохранение страницы: %s", title)

	body := r.FormValue("body")
	page := &model.Page{ID: 1, Title: title, Body: []byte(body)}
	err := pageService.SavePage(page)
	if err != nil {
		logger.Error.Printf("Ошибка сохранения страницы %s: %v", title, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info.Printf("Страница %s успешно сохранена", title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	
}

func RegisterHandlers() {
	http.HandleFunc("/view/", makeHandler(ViewHandler))
	http.HandleFunc("/edit/", makeHandler(EditHandler))
	http.HandleFunc("/save/", makeHandler(SaveHandler))
	http.HandleFunc("/registration/", makeHandler(ViewHandler))
}
