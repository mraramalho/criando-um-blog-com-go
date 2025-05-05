package main

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", blogHandler)
	router.HandleFunc("/post/", postHandler)

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	return router
}
