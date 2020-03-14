package model

// User this shit
type User struct {
	ID       *string `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	RoleID   string  `json:"role_id"`
}
