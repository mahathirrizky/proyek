package user

// CreateUserInput represents the required fields to create a new user
type CreateUserInput struct {
	Nama     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// UpdateUserInput represents the fields to update an existing user
type UpdateUserInput struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// LoginInput represents the required fields for user login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
