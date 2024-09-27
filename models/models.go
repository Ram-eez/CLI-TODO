package models

import (
	"fmt"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var AllTasks []Task

// type tasks struct {
// 	length int
// 	items  []Task
// }

func AddTask(taskName string) []Task {
	newTask := Task{
		ID:    len(AllTasks) + 1,
		Title: taskName,
		Done:  false,
	}

	AllTasks = append(AllTasks, newTask)
	fmt.Println("Task Added", newTask)
	return AllTasks
}

func listTasks() {
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

// func updateTasks(taskName string) []Task {
// 	newTask := Task{
// 		ID:    len(Tasks) + 1,
// 		Title: taskName,
// 		Done:  false,
// 	}
// }
