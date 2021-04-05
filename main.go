package main

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/connect"
	"restapi/structures"

	//"restapi/structures"
	"github.com/gorilla/mux"
)

func main() {
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user/new", NewUser).Methods("POST")
	r.HandleFunc("/user/update/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/delete/{id}", DeleteUser).Methods("DELETE")

	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	status := "success"
	var message string
	user := connect.GetUser(user_id)

	if user.Id <= 0 {
		status = "error"
		message = "User not found."
	}
	response := structures.Response{status, user, message}
	json.NewEncoder(w).Encode(response) //transporta la información al protocolo http
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	user := GetUserRequest(r)
	response := structures.Response{"success", connect.CreateUser(user), ""}
	json.NewEncoder(w).Encode(response)
}

func GetUserRequest(r *http.Request) structures.User {
	var user structures.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	user := GetUserRequest(r)
	response := structures.Response{"success", connect.UpdateUser(user_id, user), ""}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	var user structures.User
	connect.DeleteUser(user_id)
	response := structures.Response{"success", user, ""}
	json.NewEncoder(w).Encode(response)
}
