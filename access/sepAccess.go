package access

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/JayneJacobs/separateconcerns/action"
	"github.com/JayneJacobs/separateconcerns/business"
)

// Access Layer

// JSONOverHTTP Struct
type JSONOverHTTP struct {
	router  *http.ServeMux
	usrServ business.UserService
}
// NewJSONOverHTTP returns a referece to JSONOverHTTP
func NewJSONOverHTTP(usrServ business.UserService) *JSONOverHTTP {
	r := http.NewServeMux()
	joh := &JSONOverHTTP{
		router:  r,
		usrServ: usrServ,
	}
	r.HandleFunc("/register", joh.Register)
	r.HandleFunc("/user", joh.GetUser)
	return joh
}

func (j *JSONOverHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j.router.ServeHTTP(w, r)
}

// Register response to a registration request
func (j *JSONOverHTTP) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Register requires a post request", http.StatusMethodNotAllowed)
		return
	}

	params := &business.RegisterParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		http.Error(w, "Unable to read your request", http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = j.usrServ.Register(r.Context(), params)
	if err == business.ErrEmailExists {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (j *JSONOverHTTP) validateEmail(email string) error {
	if email == "" {
		return errors.New("Email must not be empty")
	}

	if !strings.ContainsRune(email, '@') {
		return errors.New("Email must include an '@' symbol")
	}

	return nil
}

// GetUser writes a response to a request
func (j *JSONOverHTTP) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "GetUser requires a get request", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	err := j.validateEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := j.usrServ.GetByEmail(r.Context(), email)
	if err == action.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}