package access

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Access Layer
type JsonOverHTTP struct {
	router  *http.ServeMux
	usrServ business.UserService
}
//  NewJsonOverHTTP 
func NewJsonOverHTTP(usrServ business.UserService) *JsonOverHTTP {
	r := http.NewServeMux()
	joh := &JsonOverHTTP{
		router:  r,
		usrServ: usrServ,
	}
	r.HandleFunc("/register", joh.Register)
	r.HandleFunc("/user", joh.GetUser)
	return joh
}

func (j *JsonOverHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j.router.ServeHTTP(w, r)
}

func (j *JsonOverHTTP) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Register requires a post request", http.StatusMethodNotAllowed)
		return
	}

	params := &RegisterParams{}
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
	if err == ErrEmailExists {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (j *JsonOverHTTP) validateEmail(email string) error {
	if email == "" {
		return errors.New("Email must not be empty")
	}

	if !strings.ContainsRune(email, '@') {
		return errors.New("Email must include an '@' symbol")
	}

	return nil
}

func (j *JsonOverHTTP) GetUser(w http.ResponseWriter, r *http.Request) {
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
	if err == ErrUserNotFound {
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