package main

import (
	storage "Issue"
	postgers "Issue/pkg/storage"
	"fmt"
	"log"
)

var db storage.Interface

func main() {

	var err error

	// Подключение к БД.
	constr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = postgers.New(constr)
	if err != nil {
		log.Fatal(err)
	}

	// Создание задачи ().
	ntask := postgers.Task{
		Title:   "Final Task (8)",
		Content: "Create tests. (8)",
	}
	id, err := db.NewTask(ntask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Новая задача создана под индексом: ", id)

	// Вывод задачи по метке.
	taskLabel, err := db.TaskByLabel("first")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(taskLabel)

	// Обновление задачи.
	utask := postgers.Task{
		ID:         4,
		AssignedID: 1,
		Title:      "Final-13",
		Content:    "tests!!!-13",
	}
	err = db.UpdateTask(utask)
	if err != nil {
		log.Fatal(err)
	}

	// Удаление задачи по ID.
	dtask := postgers.Task{ID: 10}
	err = db.DeleteTask(dtask)
	if err != nil {
		log.Fatal(err)
	}

	// Вывод всех задач.
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}
