package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {

	// POST   /users          → createUser
	r.HandleFunc("/users", h.createUser).
		Methods(http.MethodPost)

	// GET    /users/{email}  → getUserByEmail
	r.HandleFunc("/users/{email}", h.getUserByEmail).
		Methods(http.MethodGet)
}

// REFACTOR ALL THIS AND UNDERSTANDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD
// request payload for creating a user
type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// createUser handles POST /users
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user...")
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	u := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if _, err := h.svc.Create(r.Context(), u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// getUserByEmail handles GET /users/{email}
func (h *Handler) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	// 1) grab the {email} path variable
	email := mux.Vars(r)["email"]

	// 2) call the service
	u, err := h.svc.FindByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "error fetching user", http.StatusInternalServerError)
		return
	}
	if u == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// 3) return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
