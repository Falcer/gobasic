package controller

import (
	"github.com/arganaphangquestian/gobasic/model"
	"github.com/arganaphangquestian/gobasic/repository"
)

// GetAllUser handler
func GetAllUser() model.APIResponse {
	var users []model.User
	db, err := repository.OpenDB()

	defer db.Close()

	result, err := db.Query("SELECT id, email, username, password, role_id FROM users")
	for result.Next() {
		var user model.User
		_ = result.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.RoleID)
		users = append(users, user)
	}

	if err != nil {
		return model.APIResponse{
			Status:  "failed",
			Message: "Failed get users",
			Data:    nil,
		}
	}
	return model.APIResponse{
		Status:  "success",
		Message: "Get All User",
		Data:    users,
	}
}

// Register Handler
func Register(user model.User) model.APIResponse {

	db, err := repository.OpenDB()
	defer db.Close()

	stmt := `INSERT INTO users (email, username, password, role_id) VALUE (?, ?, ?, ?)`

	_, err = db.Exec(stmt, user.Email, user.Username, user.Password, "3")

	if err != nil {
		return model.APIResponse{
			Status:  "failed",
			Message: "Failed to register",
			Data:    err,
		}
	}

	return model.APIResponse{
		Status:  "success",
		Message: "Register successfully",
		Data:    user,
	}
}

// Login Handler
func Login(user model.User) model.APIResponse {
	db, err := repository.OpenDB()
	defer db.Close()

	stmt := `SELECT * FROM users WHERE username=? AND password=?`

	result := db.QueryRow(stmt, user.Username, user.Password)
	err = result.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.RoleID)

	if err != nil {
		return model.APIResponse{
			Status:  "failed",
			Message: "Failed to Login",
			Data:    err,
		}
	}

	return model.APIResponse{
		Status:  "success",
		Message: "Register successfully",
		Data:    user,
	}
}
