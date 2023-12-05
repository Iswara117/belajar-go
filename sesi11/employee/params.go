package employee

type CreateMenuRequest struct {
	Nip     string `json:"nip"`
	Name 	string `json:"name"`
	Address  string `json:"address"`
}