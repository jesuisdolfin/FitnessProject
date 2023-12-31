// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"
)

type Liftentry struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
	// must be positive
	WeightLifted string `json:"weight_lifted"`
	// must be positive
	Reps     string    `json:"reps"`
	CreateAt time.Time `json:"create_at"`
}

type UserData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// must be positive
	Weight string `json:"weight"`
	// must be positive
	Height string `json:"height"`
	// must be positive
	Age       string     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}
