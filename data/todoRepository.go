package data

import (
	"errors"
	"time"

	"../models"
)

var todoMap = map[int]models.ToDo{
	1: {
		ID:          1,
		ToDo:        "Make some exercises",
		Description: "Walk and run through the park.",
		IsDone:      false,
		CreateAt:    time.Now(),
	},
	2: {
		ID:          2,
		ToDo:        "Buy fruits and vegetables",
		Description: "They are on sale this week.",
		IsDone:      false,
		CreateAt:    time.Now(),
	},
}
var counter = 2

func GetAll() []models.ToDo {

	var todos []models.ToDo
	for _, t := range todoMap {
		todos = append(todos, t)
	}
	return todos
}

func GetByID(id int) (models.ToDo, error) {

	if _, ok := todoMap[id]; ok {
		return todoMap[id], nil
	}
	return models.ToDo{}, errors.New("Not found")
}

func Create(todo *models.ToDo) {

	counter++
	todo.CreateAt = time.Now()
	todo.ID = counter
	todoMap[counter] = *todo
}
