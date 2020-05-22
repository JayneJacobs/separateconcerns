package main

import "net/http"

func main() {
	usrstr := accesslayer.NewMemoryUserStorage()
	usrsrv := businesslayer.NewUserServiceImpl(usrstr)

	joh :=  accesslayer.NewJsonOverHttp(usrsrv)

	err := http.ListenAndServe(":"+port, joh)
	if err != nil {
		panic(err)
	}
}