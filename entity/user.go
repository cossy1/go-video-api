package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type UnixTime int

func (t UnixTime) MarshalJSON() ([]byte, error) {
	// Convert int (Unix timestamp) → time.Time → formatted string
	formatted := time.Unix(int64(t), 0).Format("02-01-2006") // dd-MM-yyyy
	return json.Marshal(formatted)
}

type User struct {
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Age       int       `json:"age" binding:"required,gte=10,lte=120"`
	Email     string    `json:"email" binding:"required,email" gorm:"unique"`
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Videos    []Video   `json:"videos" gorm:"foreignKey:UserID"`
	Password  string    `json:"password" binding:"required,min=6" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Videos    []Video   `json:"videos,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AuthorResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateUserRequest struct {
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Age       int       `json:"age" binding:"required,gte=10,lte=120"`
	UpdatedAt time.Time `json:"updatedAt"`
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
