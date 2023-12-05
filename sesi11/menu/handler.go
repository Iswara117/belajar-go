package menu

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"database/sql"
)

type Handler struct {
    // handler membutuhkan dependency
	// ke services
    services Service
}

// proses untuk menambahkan dependency kedalam handler
// jika ada dependency yg dibutuhkan, silahkan masukkan ke dalam parameter
func NewHandler(service Service) Handler {
	return Handler{
		services: service,
	}
}

// method for Create Menu
func (h Handler) Create(rw http.ResponseWriter, r *http.Request) {
    var req CreateMenuRequest
	
	// process parsing data dari client
	// dan memasukkannya kedalam variable `req`.
	// jadi variable yang dimasukkan harus alamat memory nya
    // jika bingung, ulangi lagi materi pointer
	err := json.NewDecoder(r.Body).Decode(&req)
	// fmt.Print(req)
	// println(string(err))
	if err != nil {
		// generate error response
		resp :=APIResponse{
            Status: http.StatusBadRequest,
            Message : "Create Fail",
        }
		WriteJsonResponse(rw, resp)
		return
	}
	menu := New(req.Name, req.Category, req.Desc, req.Price, req.ImageUrl)
	addId := menu.WithId(req.Id)
	// fmt.Print(addId)
	err = h.services.Save(addId)
	if err != nil {
		// generate error response
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}
    // mencoba untuk menampilkan isi dari variable request
	resp := APIResponse{
		Status: http.StatusCreated,
		Message : "Create Success",
		Payload: req,
    }

	WriteJsonResponse(rw, resp)
}



// method for GetAll Menu
func (h Handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	menus, err := h.services.repo.GetAll()
	if err != nil {
		// generate error response
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}
	// generate success response
	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: menus,
	}
	WriteJsonResponse(rw, resp)
}

// method for Get detail menu
func (h Handler) GetById(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		// generate error response
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   "id is required",
		}
		WriteJsonResponse(rw, resp)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		// generate error response
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}

	menu, err := h.services.repo.GetById(idInt)
	if err != nil {
		// generate error response
		resp := APIResponse{}

		if err == sql.ErrNoRows {
			resp = APIResponse{
				Status:  http.StatusNotFound,
				Message: "ERR NOT FOUND",
				Error:   err.Error(),
			}
		} else {
			resp = APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "ERR SERVER",
				Error:   err.Error(),
			}

		}
		WriteJsonResponse(rw, resp)
		return
	}
	// generate success response
	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: menu,
	}
	WriteJsonResponse(rw, resp)
}


func (h Handler) UpdateById(rw http.ResponseWriter, r *http.Request) {
    var req CreateMenuRequest
	
	// process parsing data dari client
	// dan memasukkannya kedalam variable `req`.
	// jadi variable yang dimasukkan harus alamat memory nya
    // jika bingung, ulangi lagi materi pointer
	err := json.NewDecoder(r.Body).Decode(&req)
	// fmt.Print(req)
	// println(string(err))
	if err != nil {
		// generate error response
		resp :=APIResponse{
            Status: http.StatusBadRequest,
            Message : "Create Fail",
        }
		WriteJsonResponse(rw, resp)
		return
	}
	menu := New(req.Name, req.Category, req.Desc, req.Price, req.ImageUrl)
	// addId := menu.WithId(req.Id)
	// fmt.Print(addId)
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		// generate error response
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}
	err = h.services.repo.UpdateById(intId, menu)
	if err != nil {
		// generate error response
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(rw, resp)
		return
	}
    // mencoba untuk menampilkan isi dari variable request
	resp := APIResponse{
		Status: http.StatusCreated,
		Message : "Update Success",
		Payload: req,
    }

	WriteJsonResponse(rw, resp)
}