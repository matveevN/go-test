package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Выводим приветственное сообщение
	fmt.Fprintln(w, "Hello, 世界")

	// Выводим ASCII-арт с изображением кота
	fmt.Fprintln(w, `
     /\_/\  
    / o o \ 
  (   "   ) 
   \~(*)~/  
    \~Y~/   
     |/|    
     |_|    
    `)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
