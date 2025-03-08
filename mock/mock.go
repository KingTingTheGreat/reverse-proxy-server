package mock

import (
	"log"
	"net/http"
)

func MockRouter(id string) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("mock ref", r.Referer())
		RenderTemplate(w, "home.html", id)
	})

	router.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(id))
	})

	return router
}

func MockServer(id string) *http.Server {
	router := MockRouter(id)

	server := http.Server{
		Handler: router,
	}

	return &server
}
