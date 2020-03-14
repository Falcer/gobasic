package route

import (
	"encoding/json"
	"net/http"

	"github.com/arganaphangquestian/gobasic/controller"
	"github.com/arganaphangquestian/gobasic/model"
)

func getAllUser(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(controller.GetAllUser())
}

func registerHandler(w http.ResponseWriter, req *http.Request) {

	var user model.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(model.APIResponse{
			Status:  "failed",
			Message: "Failed to Decode Data that you send nude 👙",
			Data:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(controller.Register(user))

}

func loginHandler(w http.ResponseWriter, req *http.Request) {

	var user model.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(model.APIResponse{
			Status:  "failed",
			Message: "Failed to Decode Data that you send nude 👙",
			Data:    err,
		})
		return
	}

	json.NewEncoder(w).Encode(controller.Login(user))
}
