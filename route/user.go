package route

import (
	"encoding/json"
	"net/http"

	"github.com/arganaphangquestian/gobasic/model"
	"github.com/arganaphangquestian/gobasic/repository"
)

func getAllUser(w http.ResponseWriter, req *http.Request) {
	var users []model.User
	db, err := repository.OpenDB()

	defer db.Close()

	result, err := db.Query("SELECT id, email, username, password, token, role_id FROM `users`")
	for result.Next() {
		var user model.User
		_ = result.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Token, &user.RoleID)
		users = append(users, user)
	}

	if err != nil {
		json.NewEncoder(w).Encode(model.APIResponse{
			Status:  "failed",
			Message: "Failed get users",
			Data:    nil,
		})
		return
	}
	json.NewEncoder(w).Encode(model.APIResponse{
		Status:  "success",
		Message: "Get All User",
		Data:    users,
	})
}

func registerHandler(w http.ResponseWriter, req *http.Request) {

	db, err := repository.OpenDB()
	defer db.Close()

	var user model.User

	err = json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(model.APIResponse{
			Status:  "failed",
			Message: "Failed to Decode Data that you send nude 👙",
			Data:    err,
		})
		return
	}

	stmt := `INSERT INTO users (email, username, password, role_id) VALUE (?, ?, ?, ?)`

	_, err = db.Exec(stmt, user.Email, user.Username, user.Password, "3")
	user.Token = nil
	user.RoleID = "3"

	if err != nil {
		json.NewEncoder(w).Encode(model.APIResponse{
			Status:  "failed",
			Message: "Failed to register",
			Data:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(model.APIResponse{
		Status:  "success",
		Message: "Register successfully",
		Data:    user,
	})
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	db, err := repository.OpenDB()
	defer db.Close()

	var userData model.User

	err = json.NewDecoder(req.Body).Decode(&userData)

	if err != nil {
		json.NewEncoder(w).Encode(model.APIResponse{
			Status:  "failed",
			Message: "Failed to Decode Data that you send nude 👙",
			Data:    err,
		})
		return
	}
	stmt := `SELECT * FROM users WHERE username=? AND password=?`

	result := db.QueryRow(stmt, userData.Username, userData.Password)
	err = result.Scan(&userData.ID, &userData.Email, &userData.Username, &userData.Password, &userData.Token, &userData.RoleID)

	if err != nil {
		json.NewEncoder(w).Encode(model.APIResponse{
			Status:  "failed",
			Message: "Failed to Login",
			Data:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(model.APIResponse{
		Status:  "success",
		Message: "Register successfully",
		Data:    userData,
	})
}
