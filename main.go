package main

import (
	"net/http"

	"github.com/JayneJacobs/separateconcerns/action"
	"github.com/JayneJacobs/separateconcerns/business"
)

func main() {
	usrstr := action.NewMemoryUserStorage()
	usrsrv := business.NewUserServiceImpl(usrstr)

	joh :=  access.NewJsonOverHttp(usrsrv)

	err := http.ListenAndServe(":"+port, joh)
	if err != nil {
		panic(err)
	}
}