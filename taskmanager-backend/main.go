package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
}

var tasks []Task

func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/tasks", getTasks).Methods("GET")
	router.HandleFunc("/api/tasks", createTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", deleteTask).Methods("DELETE")

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Adjust this in production
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS middleware
	handler := c.Handler(router)

	log.Print("server running on port 8000") // Moved before ListenAndServe
	log.Fatal(http.ListenAndServe(":8000", handler)) 
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, _ := strconv.Atoi(params["id"])
	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, _ := strconv.Atoi(params["id"])

	var updatedTask Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	for index, task := range tasks {
		if task.ID == taskID {
			tasks[index] = updatedTask
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, _ := strconv.Atoi(params["id"])

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}