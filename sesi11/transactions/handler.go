package transactions

import (
	"encoding/json"
	"strconv"
	"strings"

	// "fmt"
	"net/http"

	// "github.com/go-chi/chi"
	// "strconv"
	"github.com/go-chi/chi/v5"
	// "database/sql"
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

func (h Handler) Pay(w http.ResponseWriter, r *http.Request) {
	var req = CreateMenuRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}

	var trx = New(req)
	err := h.services.Pay(trx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			resp := APIResponse{
				Status:  http.StatusNotFound,
				Message: "DATA NOT FOUND",
				Error:   err.Error(),
			}
			WriteJsonResponse(w, resp)
			return
		}
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}
	resp := APIResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS",
	}
	WriteJsonResponse(w, resp)
}

func (h Handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	trxs, err := h.services.repo.GetAll()
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

	if len(trxs) == 0 {
		resp := APIResponse{
			Status:  http.StatusNotFound,
			Message: "DATA NOT FOUND",
			Error:   "data not found",
		}

		WriteJsonResponse(rw, resp)
		return
	}
	// generate success response
	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: trxs,
	}
	WriteJsonResponse(rw, resp)
}

func (h Handler) DetailById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   "id is required",
		}

		WriteJsonResponse(w, resp)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusBadRequest,
			Message: "ERROR BAD REQUEST",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}

	trx, err := h.services.repo.FindById(idInt)
	if err != nil {
		resp := APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "ERROR SERVER",
			Error:   err.Error(),
		}
		WriteJsonResponse(w, resp)
		return
	}

	resp := APIResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: trx,
	}
	WriteJsonResponse(w, resp)
}
