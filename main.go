package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		log.Printf("%s %s\n", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := 1234
	mux := http.NewServeMux()

	static := http.FileServer(http.Dir("./web/src/static-html"))
	mux.Handle("/", static)

	assetsDir := http.Dir("./web/dist/")
	assetsHandler := http.StripPrefix("/assets/", http.FileServer(assetsDir))
	mux.Handle("/assets/", assetsHandler)

	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 500)
		currentTime := time.Now().Format(time.RFC3339Nano)
		fmt.Fprintf(w, "Current time: %s", currentTime)
	})
	loggedMux := logRequest(mux)

	fmt.Printf("Server started on http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedMux); err != nil {
		panic(err)
	}
}

