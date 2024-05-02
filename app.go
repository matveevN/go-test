package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Отправляем заголовок Content-Type для HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Вставляем HTML-код с текстом и изображением
	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintln(w, "<p>Hello, 世界</p>")
	fmt.Fprintln(w, "<img src='/image.jpg' alt='Image' width='300' height='200'>")
	fmt.Fprintln(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/home/nikita/Dev-course/golang/image.jpg", http.FileServer(http.Dir("."))) // Указываем путь к файлу изображения
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

