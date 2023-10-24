package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type User struct {
	Name  string
	Class string
}

func saveToFile(users []User, filename string) error {
    // Mengkodekan slice users ke dalam format JSON
    jsonData, err := json.Marshal(users)
    if err != nil {
        return err
    }

    // Menyimpan data JSON ke dalam file
    err = ioutil.WriteFile(filename, jsonData, 0644)
    if err != nil {
        return err
    }

    return nil
}

func loadFromFile(filename string) ([]User, error) {
    // Membaca data JSON dari file
    jsonData, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

	// fmt.Println(jsonData)

    // Membuat slice pengguna untuk memuat data JSON yang dibaca
    var users []User

    // Mengurai data JSON ke dalam slice pengguna
    err = json.Unmarshal(jsonData, &users)
    if err != nil {
        return nil, err
    }

    return users, nil
}


func main(){

	//Data dummy Awal
	// users := []User{
	// 	User{Name: "NooB", Class: "A"},
	// 	User{Name: "Java", Class: "B"},
	// 	User{Name: "Docker", Class: "C"},
	// }
	
	user, err := loadFromFile("users.json")
	if err != nil {
		fmt.Println("Gagal membaca data dari file:", err)
	} else {
		fmt.Println("Data pengguna yang dibaca:")
		for _, user := range user {
			fmt.Printf("Name: %s, Age: %s\n", user.Name, user.Class)
		}
	}

	// optional data
	dataDariFrontend := User{
		Name:  "NooBee",
		Class: "B",
	}
	// data := []User(
	newData := append(user, dataDariFrontend)

	fmt.Println("User setelah ditambah dari frontend :", len(newData))

	err = saveToFile(newData, "users.json")
    if err != nil {
        fmt.Println("Gagal menyimpan data ke file:", err)
    } else {
        fmt.Println("Data pengguna berhasil disimpan dalam file JSON.")
    }


}