package transactions

type CreateMenuRequest struct {
	EmployeeId int      `json:"employee_id"`
	MenuId     int      `json:"menu_id"`
	Quantity   int      `json:"quantity"`
	// Total      int      `json:"total_price"`
}