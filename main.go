package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var tasks = []Task{
	{ID: 1, Name: "Learn Go"},
	{ID: 2, Name: "Build a web app"},
}

var taskID = 2

func main() {
	fmt.Println("######## Welcome to our Todolist App! ########")
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fs)
	mux.HandleFunc("/showtasks", showTasks)
	mux.HandleFunc("/addtask", addTask)
	mux.HandleFunc("/deletetask", deleteTask)
	mux.HandleFunc("/updatetask", updateTask)
	handler := enableCORS(mux)
	http.ListenAndServe(":8080", handler)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello User, welcome to our Todolist App!")
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(tasks)
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusBadRequest)
		return
	}

	var task Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	taskID++
	task.ID = taskID
	tasks = append(tasks, task)

	json.NewEncoder(writer).Encode(task)
}

func deleteTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodDelete {
		http.Error(writer, "Invalid request method", http.StatusBadRequest)
		return
	}

	idStr := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(writer, "Task \"%v\" successfully deleted", task.Name)
			return
		}
	}

	http.Error(writer, "Task not found", http.StatusNotFound)
}

func updateTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPut {
		http.Error(writer, "Invalid request method", http.StatusBadRequest)
		return
	}

	var task Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			json.NewEncoder(writer).Encode(task)
			return
		}
	}

	http.Error(writer, "Task not found", http.StatusNotFound)
}
