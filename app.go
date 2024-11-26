package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}

var tasks []Task
var taskFile = "tasks.json"

func main() {
	// Load tasks from file
	LoadTasks()
	defer SaveTasks()

	var choice int
	for{
		ShowMenu()
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Add Task")
			AddTask()
		case 2:
			fmt.Println("List Tasks")
			ListTasks()
		case 3:
			fmt.Println("Mark Task as Completed")
			LoadTaskAsCompleted()
		case 4:
			fmt.Println("Delete Task")
			DeleteTask()
		case 5:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	

	}
}

func LoadTasks(){
	// Open file
	file, err := os.Open(taskFile)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	// Read file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file", err)
		return
	}
}

func ListTasks(){
	if len(tasks) == 0 {
		fmt.Println("No tasks to display")
		return
	}
	fmt.Printf("\n%-4s | %-30s | %-40s | %-10s\n", "ID", "Name", "Description", "Status")
	fmt.Println(strings.Repeat("-", 100))
	status := ""
	for _, task := range tasks {
		status = "Not Completed"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%-4d | %-30s | %-40s | %-10s\n", task.ID, task.Name, task.Description, status)
	}
}

func ShowMenu(){
	fmt.Println("\n--- Todo List Manager ---")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mark Task as Completed")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

func AddTask(){
	var task Task
	var name, description string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Task Name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Enter Task Description: ")
	scanner.Scan()
	description = scanner.Text()

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	task.ID = id
	task.Name = name
	task.Description = description
	task.Completed = false

	tasks = append(tasks, task)
}

func LoadTaskAsCompleted(){
	var id int
	fmt.Print("Enter Task ID:")
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed")
			return
		}
	}
	fmt.Println("Task not found")
}

func DeleteTask(){
	var id int
	fmt.Print("Enter Task ID:")
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted")
			return
		}
	}
	fmt.Println("Task not found")
}

func SaveTasks(){
	// Open file
	file, err :=  os.Create(taskFile)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	// Write to file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding file", err)
		return
	}
	fmt.Println("Successfully saved tasks")

}