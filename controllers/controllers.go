package controllers

import (
	"CLI-TODO/models"
	"encoding/json"
	"fmt"
	"os"
)

func SaveToFile() {
	data, err := json.Marshal(models.AllTasks)
	if err != nil {
		fmt.Println("Error can't marshal the data", err)
		return
	}

	err = os.WriteFile("CLI-todo/data/Tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error Writing file to the data.")
		return
	}

}

func LoadTasks() {
	data, err := os.ReadFile("CLI-todo/data/Tasks.json")
	if err != nil {
		fmt.Println("Unable to read file")
		return
	}

	err = json.Unmarshal(data, &models.AllTasks)
	if err != nil {
		fmt.Println("Unable to Marshall")
	}
}

func deleteTask(ID int) {
	index := -1
	for i, task := range models.AllTasks {
		if task.ID == ID {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	models.AllTasks = append(models.AllTasks[:index], models.AllTasks[index+1:]...)
	fmt.Printf("Task with ID %d deleted successfully.\n", ID)
}

func updateTasks(taskName string, ID int, status bool) []models.Task {
	deleteTask(ID)

	newTask := models.Task{
		ID:    ID,
		Title: taskName,
		Done:  status,
	}

	models.AllTasks = append(models.AllTasks[:ID], append([]models.Task{newTask}, models.AllTasks[ID:]...)...)
	fmt.Println("Task updated successfully")
	return models.AllTasks
}
