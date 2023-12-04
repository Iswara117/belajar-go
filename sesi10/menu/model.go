package menu

import "time"

type Menu struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Desc      string    `json:"desc"`
	Price     int       `json:"price"`
	ImageUrl  string  	`json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name, category, desc string, price int, image_url string) Menu {
	return Menu{
		Name:      name,
		Category:  category,
		Desc:      desc,
		Price:     price,
		ImageUrl: image_url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// jika ingin menggunakan custom id
// Custom ID
func (m Menu) WithId(id int) Menu {
	m.Id = id
	return m
}
