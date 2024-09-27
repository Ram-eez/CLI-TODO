package main

import (
	"CLI-TODO/controllers"
	"flag"
	"fmt"
)

func main() {
	var (
		add    = flag.String("add", "", "Task name to add")
		list   = flag.Bool("list", false, "list All tasks")
		delete = flag.Int("delete", -1, "The ID for the task to delete")
		update = flag.Int("update", -1, "Task name to update")
	)

	flag.Parse()

	if *add != "" {
		fmt.Println("Adding your task")
		controllers.AddTask(*add)
	}

	if *list {
		fmt.Println("Here are all your Tasks :")
		controllers.ListTasks()
	}
	if *delete != -1 {
		fmt.Println("Deleting the task")
		controllers.DeleteTask(*delete)
	}
	if *update != -1 {
		fmt.Println("Updating your task")
		controllers.UpdateTasks(*update)
	}

}
