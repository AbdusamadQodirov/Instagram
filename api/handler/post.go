package handler

import (
	"encoding/json"
	"instagram/models"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}

	var newPost models.Post

	path := r.URL.Path
	paths := strings.Split(path, "/")
	user_id := paths[len(paths) - 1]

	newPost.UserID = user_id

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newPost); err != nil {
		log.Println("Error on creating post on handler:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	newPost.ID = uuid.New().String()

	post, err := h.Storage.GetPostRepo().CreatePost(newPost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error on create post:", err)
	}

	bytdata, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error on create post:", err)
	}

	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytdata)

}
func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error on method")
	}

	posts, err := h.Storage.GetPostRepo().GetPosts()
	if err != nil {
		w.WriteHeader(500)
		log.Println("Error on get post from db:", err)
	}

	bytdata, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(500)
		log.Println("Error on marshaling data:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytdata)

}
func (h *Handler) GetPostsById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error on method")
	}

	user_id := r.Header.Get("user_id")

	posts, err := h.Storage.GetPostRepo().GetPostsById(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytdata, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	print(bytdata)
	print(posts)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytdata)

}
func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return 
	}
	var post models.Post

	path := r.URL.Path 
	paths := strings.Split(path, "/")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	post.ID = paths[len(paths) - 1]

	postt := h.Storage.GetPostRepo().UpdatePost(post)

	bytedata, err := json.Marshal(postt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}
func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return 
	}

	path := r.URL.Path
	paths := strings.Split(path, "/")
	id := paths[len(paths) - 1]

	if err := h.Storage.GetPostRepo().DeletePost(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\nmessage : Deleted post}"))
}
