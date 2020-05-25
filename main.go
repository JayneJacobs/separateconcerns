package main

import (
	"net/http"
	"os"

	"github.com/JayneJacobs/separateconcerns/access"
	"github.com/JayneJacobs/separateconcerns/action"
	"github.com/JayneJacobs/separateconcerns/business"
)

func main() {
	usrstr := action.NewMemoryUserStorage()
	usrsrv := business.NewUserServiceImpl(usrstr)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}	
	joh :=  access.NewJSONOverHTTP(usrsrv)

	err := http.ListenAndServe(":"+port, joh)
	if err != nil {
		panic(err)
	}
}