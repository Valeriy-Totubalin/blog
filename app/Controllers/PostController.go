package Controllers

import (
	"../Services"
	"fmt"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/post-new.html", "templates/head.html", "templates/nav-bar.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "post-new", nil)
}

func WriteHandler(w http.ResponseWriter, r *http.Request)  {
	Services.AddPost(r)
	http.Redirect(w, r, "/post/list", 302)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/post-list.html", "templates/head.html", "templates/nav-bar.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	posts := Services.GetPostByPages(45)
	t.ExecuteTemplate(w, "post-list", posts)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	post := Services.GetEditPost(r)
	t, err := template.ParseFiles("templates/post-edit.html", "templates/head.html", "templates/nav-bar.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "post-edit", post)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request)  {
	Services.UpdatePost(r)
	http.Redirect(w, r, "/post/list", 302)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request)  {
	Services.DeletePost(r)
	http.Redirect(w, r, "/post/list", 302)
}
