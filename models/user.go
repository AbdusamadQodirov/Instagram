package models

type User struct {
    ID       string `json:"id"`
    Email    string `json:"email"`
    Fullname string `json:"fullname"`
    Username string `json:"username"`
    Password string `json:"password"`
}


