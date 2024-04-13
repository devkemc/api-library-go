package routes

import (
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
	"github.com/devkemc/api-library-go/pkg/web/response"
	"github.com/gorilla/mux"
	"net/http"
)

func RouterInit() {
	postgresConn, errConnection := data.ConnectDB()
	if errConnection != nil {
		panic(errConnection)
	}
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	responseJson := response.NewJsonResponse()
	InitBooksRoutes(api, postgresConn, responseJson)
	InitAuthorsRoutes(api, postgresConn, responseJson)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	err := http.ListenAndServe(":8081", r)
	if err != nil {

		panic("error when running server")
	}
}
