package model

// User this shit
type User struct {
	ID       *string `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Token    *string `json:"token"`
	RoleID   string  `json:"role_id"`
}
