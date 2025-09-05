package entity

type User struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Age       int     `json:"age" binding:"required,gte=10,lte=120"`
	Email     string  `json:"email" binding:"required,email" gorm:"unique"`
	ID        uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Videos    []Video `json:"videos" gorm:"foreignKey:UserID"`
	Password  string  `json:"password" binding:"required,min=6" gorm:"not null"`
	CreatedAt int     `json:"createdAt"`
}

type UserResponse struct {
	ID        uint64  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Age       int     `json:"age"`
	Email     string  `json:"email"`
	Videos    []Video `json:"videos,omitempty"`
	CreatedAt int     `json:"createdAt"`
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
	}
}
