package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../data"
	"../models"
)

// GetAll returns all todos
func GetAll(w http.ResponseWriter, r *http.Request) {

	todos := data.GetAll()
	w.Header().Set("Content-Type", "application/json")
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(todosJSON)

}

// GetByID returns a todo entity using ID
func GetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(vars["id"])
	todo, err := data.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{}"))
	} else {
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(todoJSON)
	}
}

// Create saves and returns a new todo entity
func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		panic(err)
	}
	data.Create(&todo)

	todoJSON, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(todoJSON)
}
