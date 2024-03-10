package model

import "time"

type User struct {
	Id        int
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Repeat    string `json:"repeat"`
	Role      string `json:"role"`
	ExpiresAt time.Time
	IsAuth    bool
}
