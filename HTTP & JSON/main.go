package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var users = []User{}

type api struct {
	addr string
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Implement basic routing
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Home Page"))
			return
		case "/users":
			w.Write([]byte("Users page"))

		}
	default:
		w.Write([]byte("404 Page"))
	}
}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Encode the user slice to JSON
	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body to User struct
	var payload User

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	// users = append(users, u)

	// Using out custom insert method to add some validations
	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	// Input Validation
	if u.FirstName == "" {
		return errors.New("Firstname is required")
	}
	if u.LastName == "" {
		return errors.New("Lastname is required")
	}

	// Storage Validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already exists")
		}
	}

	users = append(users, u)
	return nil
}

func main() {
	api := &api{addr: ":8080"} // Creating a inapitance of our api handler

	// err := http.ListenAndServe(api.addr, api)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// We can add additional config using this way
	// srv := &http.Server{
	// 	Addr:    api.addr,
	// 	Handler: api,
	// }

	// srv.ListenAndServe()

	// Initialize the ServeMux for handing routing
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
