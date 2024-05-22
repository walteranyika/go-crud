package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}

type Measurement struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	Weight    float64   `json:"weight"`
	Height    float64   `json:"height"`
	BodyFat   float64   `json:"body_fat"`
	CreatedAt time.Time `json:"created_at"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
