package handler

import (
	"encoding/json"
	"instagram/models"
	"instagram/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}

	var newUser models.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newUser); err != nil {
		w.WriteHeader(500)
		return
	}

	newUser.ID = uuid.New().String()

	user, err := h.Storage.GetUserRepo().CreateUser(newUser)
	if err != nil {
		log.Println("Error on save data", err)
		w.WriteHeader(500)

	}

	bytdata, err := json.Marshal(user)
	if err != nil {
		log.Println("Error on change format", err)
		w.WriteHeader(500)
	}

	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytdata)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Error on convert data limit:", err)
		w.WriteHeader(500)
	}
	if limit == "" || limit == "0" {
		limitInt = 0
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error on convert data page:", err)
		w.WriteHeader(500)
		return
	}
	if page == "" || page == "0" {
		pageInt = 0
	}

	users, err := h.Storage.GetUserRepo().GetUsers(pkg.Limit(limitInt), pkg.Page(pageInt))
	if err != nil {
		log.Println("Error on get data from database:", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error on method")
		return
	}
	//path on url
	path := r.URL.Path

	// each path save separately
	paths := strings.Split(path, "/")

	id := paths[len(paths)-1]

	user, err := h.Storage.GetUserRepo().GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error on formatting data to json:", err)
		return
	}

	w.Write(jsonData)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	paths := strings.Split(path, "/")

	var user models.User
	decoder := json.NewDecoder(r.Body)

	user.ID = paths[len(paths)-1]

	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userr := h.Storage.GetUserRepo().UpdateUser(user)

	bytedata, err := json.Marshal(userr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytedata)
}

func (h *Handler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path := r.URL.Path

	paths := strings.Split(path, "/")
	id := paths[len(paths)-1]

	err := h.Storage.GetUserRepo().DeleteUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("{\nmessage : Deleted user\n}"))
}
