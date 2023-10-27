package main

import (
	"log"
	"net/http"
	"weather-apps/app"
)

func main() {
    /* 
        bebas mau dibuat berapa port nya
        example : 4444,4000,8080,8000, ...etc
        yang penting harus belum pernah digunakan
    */
	port := ":4444"

	http.HandleFunc("/weather", app.GetCurrentWeather)
    // tampilkan log sebagai penanda bahwa server
    // telah running
	log.Println("server running at port", port)

    // proses running server
	http.ListenAndServe(port, nil)
}
