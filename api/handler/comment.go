package handler

import (
	"encoding/json"
	"instagram/models"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var comment models.Comment

	comment.ID = uuid.New().String()

	user_id := r.URL.Query().Get("user_id")
	post_id := r.URL.Query().Get("post_id")

	comment.UserID = user_id
	comment.PostID = post_id
	
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&comment); err != nil {
		log.Println("Error on decoding")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	commentt, err := h.Storage.GetCommentRepo().CreateComment(comment)
	if err != nil {
		log.Println("Error creating data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytedata, err := json.Marshal(commentt)
	if err != nil {
		log.Println("Error on marshaling data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)

}

func (h *Handler) GetCommentByPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	paths := strings.Split(path, "/")
	id := paths[len(paths)-1]

	comments, err := h.Storage.GetCommentRepo().GetCommentByPost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytedata, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, "Error on marshaling data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var comment models.Comment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&comment); err != nil {
		http.Error(w, "Error on decoding data", http.StatusInternalServerError)
		return
	}

	commentt := h.Storage.GetCommentRepo().UpdateComment(comment)

	bytedata, err := json.Marshal(commentt)
	if err != nil {
		http.Error(w, "Error on marshaling data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	paths := strings.Split(path, "/")

	id := paths[len(paths)-1]

	if err := h.Storage.GetCommentRepo().DeleteComment(id); err != nil {
		http.Error(w, "Error on deleting data from database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\nmessage : Deleted comment}"))

}
