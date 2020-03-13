package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arganaphangquestian/gobasic/middleware"
	"github.com/arganaphangquestian/gobasic/model"
	"github.com/gorilla/mux"
)

func handleBase(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(model.APIResponse{
		Status:  "Success",
		Message: "Yeay 🔥",
	})
}

// Init this shit
func Init() {
	const port string = ":8080"
	r := mux.NewRouter()
	// Middleware
	r.Use(middleware.JSON())
	r.Use(middleware.CORS())
	// Routing
	r.HandleFunc("/", handleBase).Methods("GET")
	r.HandleFunc("/users", getAllUser).Methods("GET")
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	fmt.Printf("Yeay, this shit at http://127.0.0.1%s 🔥\n", port)
	http.ListenAndServe(port, r)
}
