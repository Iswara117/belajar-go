package utility

import (
	"math/rand"
	"time"
)

// fungsi untuk generate random berdasarkan jarak angka minimal dan maksimal
func GenerateRandomNumber(min int, max int) (number int) {
    // proses untuk melakukan seeding data agar bisa membuat data random
	rand.Seed(time.Now().UnixNano())

    // proses melakukan random data berdasarkan angka minimum dan maksimum.
	return rand.Intn(max-min) + min
}