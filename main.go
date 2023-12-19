// main.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var secretKey = []byte("your-secret-key")

var (
	tasksCreated = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tasks_created_total",
			Help: "Total number of tasks created",
		},
		[]string{},
	)
)

func init() {
	prometheus.MustRegister(tasksCreated)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/token", tokenHandler)
	http.Handle("/metrics", promhttp.Handler())

	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Error generating JWT token.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, signedToken)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handleBadRequest(w, "Method not allowed.")
		return
	}

	fmt.Fprint(w, "Hello, this is the home page!")

	if !validateToken(r) {
		http.Error(w, "Unauthorized. Invalid or missing token.", http.StatusUnauthorized)
		return
	}

	// Increment the tasksCreated metric
	tasksCreated.WithLabelValues().Inc()
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleBadRequest(w, "Method not allowed.")
		return
	}

	if err := r.ParseForm(); err != nil {
		handleBadRequest(w, "Error parsing form data.")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Received POST request with username: %s, password: %s\n", username, password)

	if !validateToken(r) {
		http.Error(w, "Unauthorized. Invalid or missing token.", http.StatusUnauthorized)
		return
	}

	// Increment the tasksCreated metric
	tasksCreated.WithLabelValues().Inc()

	fmt.Fprint(w, "POST request received successfully.")
}

func validateToken(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	return err == nil && token.Valid
}
