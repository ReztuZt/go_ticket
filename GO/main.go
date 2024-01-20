package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Pesanan struct {
	Nama     string `json:"nama"`
	Destinasi string `json:"destinasi"`
	Status   string `json:"status"`
	Jumlah   int    `json:"jumlah"`
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/pesan-tiket", PesanTiketHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func PesanTiketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode HTTP tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var pesanan Pesanan
	err := decoder.Decode(&pesanan)
	if err != nil {
		http.Error(w, "Gagal membaca body request", http.StatusBadRequest)
		return
	}


	// Contoh respons JSON
	response := map[string]interface{}{
		"success": true,
		"message": "Pesanan berhasil diterima",
		"data":    pesanan,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
