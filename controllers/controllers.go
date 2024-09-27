package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
	Done  bool   `json:"Done"`
}

var AllTasks []Task

func LoadTasks() {

	data, err := os.ReadFile("data/Tasks.json")
	if err != nil {
		fmt.Println("Unable to read file")
		return
	}

	if len(data) != 0 {
		err = json.Unmarshal(data, &AllTasks)
		if err != nil {
			fmt.Println("Unable to Marshall")
		}
	}

}

func SaveToFile() {

	data, err := json.Marshal(AllTasks)
	if err != nil {
		fmt.Println("Error can't marshal the data:", err)
		return
	}

	err = os.WriteFile("data/Tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error Writing file to the data.")
		return
	}

}

func DeleteTask(ID int) {
	LoadTasks()
	index := -1
	for i, task := range AllTasks {
		if task.ID == ID {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	AllTasks = append(AllTasks[:index], AllTasks[index+1:]...)
	fmt.Printf("Task with ID %d deleted successfully.\n", ID)
	SaveToFile()
}

func UpdateTasks(taskID int) {

	LoadTasks()
	index := -1
	for i, task := range AllTasks {
		if task.ID == taskID {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the new Task name :")
	scanner.Scan()
	taskName := scanner.Text()

	fmt.Println("Enter the updated task status (true or false)")
	scanner.Scan()
	status, _ := strconv.ParseBool(scanner.Text())

	updatedTask := Task{
		ID:    taskID,
		Title: taskName,
		Done:  status,
	}

	AllTasks[index] = updatedTask
	SaveToFile()
}

func AddTask(taskName string) {

	LoadTasks()
	newTask := Task{
		ID:    len(AllTasks) + 1,
		Title: taskName,
		Done:  false,
	}

	AllTasks = append(AllTasks, newTask)
	fmt.Printf("%+v\n", newTask)
	SaveToFile()
}

func ListTasks() {

	LoadTasks()
	if len(AllTasks) == 0 {
		fmt.Println("No Tasks Available")
	}
	for _, task := range AllTasks {
		status := "[X]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%d: %s %s \n", task.ID, status, task.Title)
	}

}

// func updateTasks(taskName string, ID int, status bool) []models.Task {
// 	deleteTask(ID)

// 	newTask := models.Task{
// 		ID:    ID,
// 		Title: taskName,
// 		Done:  status,
// 	}

// 	models.AllTasks = append(models.AllTasks[:ID], append([]models.Task{newTask}, models.AllTasks[ID:]...)...)
// 	fmt.Println("Task updated successfully")
// 	return models.AllTasks
// }
