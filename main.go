package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	HelloMsg string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	m := Message{"Hello world from Golang"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}

func main() {

	fmt.Println("Server is starting...")

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is listening at 1234")

	http.ListenAndServe(":1234", nil)

}
