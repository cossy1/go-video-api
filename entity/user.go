package entity

import (
	"encoding/json"
	"time"
)

type UnixTime int

func (t UnixTime) MarshalJSON() ([]byte, error) {
	// Convert int (Unix timestamp) → time.Time → formatted string
	formatted := time.Unix(int64(t), 0).Format("02-01-2006") // dd-MM-yyyy
	return json.Marshal(formatted)
}

type User struct {
	FirstName string   `json:"firstName" binding:"required"`
	LastName  string   `json:"lastName" binding:"required"`
	Age       int      `json:"age" binding:"required,gte=10,lte=120"`
	Email     string   `json:"email" binding:"required,email" gorm:"unique"`
	ID        uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Videos    []Video  `json:"videos" gorm:"foreignKey:UserID"`
	Password  string   `json:"password" binding:"required,min=6" gorm:"not null"`
	CreatedAt UnixTime `json:"createdAt"`
	UpdatedAt UnixTime `json:"updatedAt"`
}

type UserResponse struct {
	ID        uint64   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Age       int      `json:"age"`
	Email     string   `json:"email"`
	Videos    []Video  `json:"videos,omitempty"`
	CreatedAt UnixTime `json:"createdAt"`
	UpdatedAt UnixTime `json:"updatedAt"`
}

type UpdateUserRequest struct {
	FirstName string   `json:"firstName" binding:"required"`
	LastName  string   `json:"lastName" binding:"required"`
	Age       int      `json:"age" binding:"required,gte=10,lte=120"`
	Email     string   `json:"email" binding:"required,email" gorm:"unique"`
	UpdatedAt UnixTime `json:"updatedAt"`
}

func ToUserResponse(user User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Email:     user.Email,
		Videos:    user.Videos,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
