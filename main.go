package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	/*
		db, err := NewDB()

		if err != nil {
			panic(err)
		}
	*/

	r := mux.NewRouter()
	r.HandleFunc("/send/{user}", SendHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	s := vars["user"]

	fmt.Println(s)

	w.Write([]byte(s))
}

func CreateUserHandler(_ http.ResponseWriter, _ *http.Request) {

}
