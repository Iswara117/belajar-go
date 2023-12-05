package menu

type CreateMenuRequest struct {
	Id 		 int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
}