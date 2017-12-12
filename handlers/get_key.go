package handlers

import (
	"github.com/vladborsh/go-rest/storage"
	"net/http"
)

func GetKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Missing key name in query string ", http.StatusBadRequest)
			return
		}
		val, err := db.Get(key)
		if err != nil {
			http.Error(w, "Missing key name in query string ", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}
