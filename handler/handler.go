package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// untuk menampilkan log aktivitas error
	log.Printf(r.URL.Path)
	// agar tidak semua link asal bisa masuk ke halaman home
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Is Happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := []entity.Product{
		{ID: 01, Name: "Mobilio", Price: 200000000, Stock: 4},
		{ID: 02, Name: "Xpander", Price: 277000000, Stock: 2},
		{ID: 03, Name: "Avanza", Price: 267000000, Stock: 11},
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Kaleeeeeeeeeeeeeemmm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, saya sedang belajar golang web"))
}

func Mariobros(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario Bros from nintendo win"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	data := map[string]interface{}{
		"content": idNumb,
	}
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini Adalah GET"))
	case "POST":
		w.Write([]byte("Ini Adalah POST"))
	default:
		http.Error(w, "Error is Happenig, keep calm", http.StatusBadRequest)
	}
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}
		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "keep calm, is easy", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}
