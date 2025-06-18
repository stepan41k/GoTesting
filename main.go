package main

import (
	"net/http"

	"github.com/go-chi/chi"
	sum "github.com/stepan41k/lessonPostgres/task-10/sum"
	display "github.com/stepan41k/lessonPostgres/task-5/display"
)

func main() {

	router := chi.NewRouter()

	router.Post("/user", display.DisplayName)
	router.Post("/positive-sum", sum.MultiplyInteger())

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}