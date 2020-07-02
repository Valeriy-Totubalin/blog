package main

import (
	"./app/Controllers"
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("Server started")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets/"))))
	http.HandleFunc("/", Controllers.IndexHandler)
	http.HandleFunc("/post/new", Controllers.NewHandler)
	http.HandleFunc("/post/write", Controllers.WriteHandler)
	http.HandleFunc("/post/list", Controllers.ListHandler)
	http.HandleFunc("/post/edit", Controllers.EditHandler)
	http.HandleFunc("/post/update", Controllers.UpdateHandler)
	http.HandleFunc("/post/delete", Controllers.DeleteHandler)

	http.ListenAndServe(":80", nil)
}
