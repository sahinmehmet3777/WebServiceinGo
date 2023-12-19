// handlers.go

package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var workerPool = make(chan struct{}, 10)
var wg sync.WaitGroup

var handlerTasksCreated = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "tasks_created_total",
		Help: "Total number of tasks created",
	},
	[]string{},
)

func init() {
	prometheus.MustRegister(handlerTasksCreated)
}

// createTaskHandler handles POST requests to create a new task concurrently
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		handleBadRequest(w, "Error parsing form data.")
		return
	}

	// Extract task data from the form
	title := r.FormValue("title")
	description := r.FormValue("description")
	status := r.FormValue("status")

	// Validate task fields
	err = validateTaskFields(title, description, status)
	if err != nil {
		handleBadRequest(w, fmt.Sprintf("Invalid input: %s", err))
		return
	}

	// Limit the number of concurrent tasks being processed
	workerPool <- struct{}{}
	defer func() { <-workerPool }()

	// Create a new task
	taskID, err := createTask(title, description, status)
	if err != nil {
		handleInternalError(w, fmt.Sprintf("Error creating task: %s", err))
		return
	}

	// Increment the tasksCreated metric with the "status" label
	handlerTasksCreated.WithLabelValues(status).Inc()

	// Respond with the ID of the created task
	fmt.Fprintf(w, "Task created with ID: %d", taskID)
}

func handleInternalError(w http.ResponseWriter, s string) {
	http.Error(w, s, http.StatusInternalServerError)
}

func handleBadRequest(w http.ResponseWriter, s string) {
	http.Error(w, s, http.StatusBadRequest)
}

// ... (diğer fonksiyonlar aynen devam eder)
