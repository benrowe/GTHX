package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 1234
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	fmt.Println(fmt.Sprintf("Server started on http://localhost:%d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
