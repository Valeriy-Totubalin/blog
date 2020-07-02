package Services

import (
	"../Models"
	"fmt"
	"net/http"
	"strconv"
)

var table = "posts"

func AddPost(r *http.Request)  {
	title := strconv.Quote(r.FormValue("title"))
	content := strconv.Quote(r.FormValue("content"))
	post := Models.Post{Title: title, Content: content}
	if isValid(post) {
		NewPost(post)
	}
}

func GetEditPost(r *http.Request) Models.Post {
	id := strconv.Quote(r.FormValue("id"))
	db := GetDB()
	rows := db.query("select id, title, content from posts where id=" + id + "limit 1")
	posts := []Models.Post{}
	for rows.Next() {
		p := Models.Post{}
		err := rows.Scan(&p.Id, &p.Title, &p.Content)
		if err != nil{
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	return posts[0]
}

func NewPost(post Models.Post) {
	db := GetDB()
	db.insert(table, "0 , " + post.Title + ", " + post.Content )
}

//toDo поправить вывод по страинцам
func GetPostByPages(elementCount int) []Models.Post  {
	db := GetDB()
	rows := db.query("select * from posts")
	posts := []Models.Post{}
	for rows.Next(){
		p := Models.Post{}
		err := rows.Scan(&p.Id, &p.Title, &p.Content)
		if err != nil{
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	return posts
}

func isValid(post Models.Post) bool {
	if post.Title == "" {
		return false
	}
	if post.Content == "" {
		return false
	}
	return true
}

func UpdatePost(r *http.Request) {
	id := strconv.Quote(r.FormValue("id"))
	title := strconv.Quote(r.FormValue("title"))
	content := strconv.Quote(r.FormValue("content"))
	post := Models.Post{Title: title, Content: content}
	if isValid(post) {
		db := GetDB()
		db.update(table, "title=" + title + ", content=" + content + " where id=" + id)
	}
}

func DeletePost(r *http.Request) {
	id := strconv.Quote(r.FormValue("id"))
	db := GetDB()
	db.delete(table, "id=" + id)
}