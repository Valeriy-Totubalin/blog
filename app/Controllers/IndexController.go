package Controllers

import (
	"fmt"
	"net/http"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	// парсятся шаблоны и записываются в t, если есть ошибка, то в лежит в err
	t, err := template.ParseFiles("templates/index.html", "templates/head.html", "templates/nav-bar.html")
	// Если есть ошибка, то вывести ее на экран
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	// Подключение шаблона
	t.ExecuteTemplate(w, "index", nil)
}
