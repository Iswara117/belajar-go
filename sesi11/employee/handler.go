package employee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	// "database/sql"
	// "strconv"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-delve/delve/service"
)

type Handler struct {
	services Service
}

func NewHandler(service Service ) Handler{
	return Handler{
		services: service, 
	}
}


func (h Handler) Create(rw http.ResponseWriter, r *http.Request){
	var req CreateMenuRequest

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
	employee := New(req.Name, req.Nip, req.Address)
	fmt.Print("masuk")
	// addId := menu.WithId(req.Id)
	// fmt.Print(addId)
	err = h.services.Save(employee)
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

// func (h Handler) GetAll(rw http.ResponseWriter, r *http.Request) {
// 	menus, err := h.services.repo.GetAll()
// 	if err != nil {
// 		// generate error response
// 		resp := APIResponse{
// 			Status:  http.StatusBadRequest,
// 			Message: "ERR BAD REQUEST",
// 			Error:   err.Error(),
// 		}
// 		WriteJsonResponse(rw, resp)
// 		return
// 	}
// 	// generate success response
// 	resp := APIResponse{
// 		Status:  http.StatusOK,
// 		Message: "SUCCESS",
// 		Payload: menus,
// 	}
// 	WriteJsonResponse(rw, resp)
// }

func (h Handler) GetAll( rw http.ResponseWriter, r *http.Request) {
	employees, err := h.services.repo.GetAll()
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

	if employees == nil {
		// generate error response
		resp := APIResponse{
			Status:  http.StatusNotFound,
			Message: "DATA NOT FOUND",
			Error:   "data not found",
		}
		WriteJsonResponse(rw, resp)
		return
	}

	// fmt.Print(employees)

	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: employees,
	}
	WriteJsonResponse(rw, resp)
	// return
}

func (h Handler) DeleteById(rw http.ResponseWriter, r *http.Request) {
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
	_, err = h.services.repo.GetById(idInt)
	// employee,err = h.services.repo.GetById(idInt)
	// if employee.Id == 0 {
	// 	// generate error response

	// 	resp :=APIResponse{
	// 		Status: http.StatusNotFound,
	// 		Message : "DATA NOT FOUND",
	// 		Error: "data not found",
	// 	}
	// 	WriteJsonResponse(rw, resp)
	// 	return
	// }
	if err != nil {
		// generate error response
		if err.Error() == "sql: no rows in result set"{
			resp :=APIResponse{
						Status: http.StatusNotFound,
						Message : "DATA NOT FOUND",
						Error: "data not found",
					}
					WriteJsonResponse(rw, resp)
					return
		} 
		resp :=APIResponse{
            Status: http.StatusBadRequest,
            Message : "GET DATA FAILED",
			Error: err.Error(),
        }
		WriteJsonResponse(rw, resp)
		return
	}

	
	// fmt.Print(employee)
	// if employee == nil {
	// 	// generate error response
	// 	resp := APIResponse{
	// 		Status:  http.StatusNotFound,
	// 		Message: "DATA NOT FOUND",
	// 		Error:   "data not found",
	// 	}
	// 	WriteJsonResponse(rw, resp)
	// 	return
	// }
		// generate error response

	// 	resp :=APIResponse{
    //         Status: http.StatusBadRequest,
    //         Message : "Delete Fail",
    //     }
	// 	WriteJsonResponse(rw, resp)
	// 	return
	// }
	err = h.services.repo.DeleteById(idInt)
	if err != nil {
		// generate error response

		resp :=APIResponse{
            Status: http.StatusBadRequest,
            Message : "Delete Fail",
        }
		WriteJsonResponse(rw, resp)
		return
	}
	// generate success response
	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "Delete Success",
	}
	WriteJsonResponse(rw, resp)
}