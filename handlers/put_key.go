package handlers

import (
	"github.com/vladborsh/go-rest/storage"
	"io/ioutil"
	"net/http"
)

func PutKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Missing key name in path", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}
		if err := db.Put(key, val); err != nil {
			http.Error(w, "Error setting value in DB", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
