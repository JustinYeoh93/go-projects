package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type userController struct {
	database *database
}

func newUserController(database *database) *userController {
	return &userController{
		database: database,
	}
}

func (ctrl *userController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	ctrl.database.db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func (ctrl *userController) newUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	ctrl.database.db.Create(&User{Name: name, Email: email})
	fmt.Fprint(w, "new user created")
}

func (ctrl *userController) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	ctrl.database.db.Where("name = ?", name).Find(&user)

	user.Email = email

	ctrl.database.db.Save(&user)
	fmt.Fprint(w, "user updated")
}

func (ctrl *userController) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	ctrl.database.db.Where("name = ?", name).Find(&user)
	ctrl.database.db.Delete(&user)

	fmt.Fprint(w, "user deleted")
}
