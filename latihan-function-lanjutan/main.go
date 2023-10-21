package main

import "fmt"


func countSiswa(v []float32)  int{
	count := 0

	for i := 0; i<len(v); i++{
		count += 1
	}

	return(count)
}

func getAverage(v []float32, countSiswa int)  float32{
	var nilai float32
	var average float32

	for i := 0; i<len(v); i++{
		nilai += v[i]
	}

	average = nilai /float32(countSiswa) 
	return(average)
}

func main()  {
	nilaiSiswa := []float32{85.5, 92.0, 78.5, 90.0, 88.5}

	jumlahSiswa := countSiswa(nilaiSiswa)

	fmt.Println("Jumlah nilai siswa dalam kelas:",jumlahSiswa)

	average := getAverage(nilaiSiswa, countSiswa(nilaiSiswa))
	fmt.Printf("Nilai rata-rata siswa dalam kelas: %.2f\n",average)
}