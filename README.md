# **Todo List Manager**

## **Overview**

This is a hobby project designed to practice and improve Go (Golang) skills. It is an early-stage, basic implementation of a **Todo List Management System**, focusing on fundamental backend concepts like file handling, JSON serialization, and console-based user interaction.

The application allows users to:
- Add tasks with a name and description.
- View all tasks in a formatted list.
- Mark tasks as completed.
- Delete tasks.
- Save and load tasks to/from a JSON file for persistence.

The project aims to provide a simple, functional starting point for building more advanced task management systems.

---

## **Features**

- **Add Tasks**: Add tasks with an ID, name, description, and status.
- **List Tasks**: View all tasks in a clean and aligned table format.
- **Mark as Completed**: Update the status of a task to "Completed."
- **Delete Tasks**: Remove tasks by their ID.
- **Persistent Storage**: Tasks are saved to a `tasks.json` file and loaded when the application restarts.

---

## **Getting Started**

### Prerequisites

- [Go Programming Language](https://golang.org/) (1.16 or later)

### Running the Application

1. Clone or download this repository.
2. Navigate to the project directory in your terminal.
3. Build and run the project:
   ```bash
   go run app.go
