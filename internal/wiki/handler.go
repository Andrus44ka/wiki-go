package wiki

import (
	"gowiki/internal/logger"
	"html/template"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	logger.Info.Printf("Получен запрос на просмотр страницы: %v", title)
	p, err := LoadPageFromFile(title)
	if err != nil {
		logger.Warn.Printf("Страница %s не найдена, перенаправление на /edit", title)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	logger.Info.Printf("Страница %s успешно загружена", title)
	renderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	logger.Info.Printf("Получен запрос на редактирование страницы: %s", title)

	p, err := LoadPageFromFile(title)
	if err != nil {
		logger.Warn.Printf("Страница %s не найдена, создается новая", title)
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	logger.Info.Printf("Получен запрос на сохранение страницы: %s", title)

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := SavePageToFile(p)
	if err != nil {
		logger.Warn.Printf("Ошибка сохранения страницы %s: %v ", title, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info.Printf("Страница %s успешно сохранена", title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func RegisterHandlers() {
	http.HandleFunc("/view/", makeHandler(ViewHandler))
	http.HandleFunc("/edit/", makeHandler(EditHandler))
	http.HandleFunc("/save/", makeHandler(SaveHandler))
}
