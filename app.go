package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "<h1>Hello, 世界</h1>")

	
	fmt.Fprintln(w, `
	<style>
		.cat {
			color: #3E721D; /* Цвет кота */
			font-size: 24px; /* Размер шрифта */
			font-family: Arial, sans-serif; /* Шрифт */
		}
	</style>
	<pre class="cat">
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
