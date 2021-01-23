package dto

// Login credential are required
type LoginCredentials struct {
	User     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
