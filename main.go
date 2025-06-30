package main

import "net/http"

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World! Pratik Sharma AWS from self hosted runner"))
	})
	server := &http.Server{
		Addr:    ":8070",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
