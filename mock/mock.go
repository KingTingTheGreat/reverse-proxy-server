package mock

import (
	"net/http"
)

func Router(id string) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(id))
	})

	return router
}

func Server(id string) *http.Server {
	router := Router(id)

	server := http.Server{
		Handler: router,
	}

	return &server
}
