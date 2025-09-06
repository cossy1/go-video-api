package entity

type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Age       int    `json:"age" binding:"required,gte=10,lte=120"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	ID        uint64 `json:"id"`
	CreatedAt UnixTime    `json:"createdAt"`
	UpdatedAt UnixTime    `json:"updatedAt"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
