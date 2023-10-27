package app

import (
	"encoding/json"
	// "math/rand"
	"net/http"
	"time"
	"log"

    "weather-apps/utility"
)

/*
fungsi ini bertugas sebagai handler untuk endpoint GET /weather
yang mana akan mengembalikan/menampilkan nilai random berdasarkan
kriteria yang telah dibuat sebelumnya.
*/
func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Agent")
	w.WriteHeader(http.StatusOK)

    // setup default value
	humidity := utility.GenerateRandomNumber(0,100)
	temperature := utility.GenerateRandomNumber(-10,50)
	wind := utility.GenerateRandomNumber(0,20)
	rain := utility.GenerateRandomNumber(0,250)

    // setup response
    // harus sesuai dengan API Spesifikasi yang telah dibuat
	response := map[string]interface{}{
		"humidity":     humidity,
		"temperature":  temperature,
		"wind":         wind,
		"rain":         rain,
		"last_checked": time.Now(),
	}

    // jangan lupa untuk log endpointnya
    // disini kita akan buat log untuk method, type, dan path nya.
    log.Printf("type=%v method=%v path=%v", "[INFO]", r.Method, r.URL.Path)

    // proses kirim data ke client
	json.NewEncoder(w).Encode(response)
}