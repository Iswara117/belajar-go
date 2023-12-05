package transactions

import "time"


type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Menu struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Transactions struct {
	Id        	int     `json:"id"`
	EmployeeId int      `json:"employee_id"`
	MenuId     int      `json:"menu_id"`
	Quantity   int      `json:"quantity"`
	Total      int      `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Employee Employee `json:"employee"`
	Menu     Menu     `json:"menu"`
}

func New(data CreateMenuRequest) Transactions {
	return Transactions{
		EmployeeId:   data.EmployeeId,
		MenuId:  	  data.MenuId,
		Quantity: 	  data.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *Transactions) SetTotal(price int) {
	t.Total = t.Quantity * price
}