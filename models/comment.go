package models

type Comment struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Msg    string `json:"msg"`
    UserID string `json:"user_id"`
    PostID string `json:"post_id"`
}