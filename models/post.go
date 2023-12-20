package models

type Post struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Msg    string `json:"msg"`
    UserID string `json:"user_id"`
}