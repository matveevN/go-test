package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Выводим приветственное сообщение
	fmt.Fprintln(w, "<h1>Hello, 世界</h1>")

	// Выводим ASCII-арт с изображением кота с применением CSS-стилей
	fmt.Fprintln(w, `
	<style>
		.cat {
			color: #3E721D; /* Цвет тела кота */
		}
		.cat_eyes {
			color: #FFFF00; /* Цвет глаз кота */
		}
	</style>
	<pre>
	     /\_/\  
	    / o o \ 
	  (   "   ) 
	   \~(*)~/  
	    \~Y~/   
	     |/|    
	     |_|    
	</pre>
	`)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
