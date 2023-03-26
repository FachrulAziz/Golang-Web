package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	// mux merupakan variable agar bisa menampilkan di localhost
	mux := http.NewServeMux()

	// cara mendaftarkan halaman web
	mux.HandleFunc("/", handler.HomeHandler) // <= tanda "/" artinya adalah root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.Mariobros)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) { // <= mendaftarkan sekaligus mendeklarasikan sebuah variable
		w.Write([]byte("Profile Page"))
	})

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
