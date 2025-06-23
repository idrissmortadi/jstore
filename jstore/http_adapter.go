package jstore

import (
	"encoding/json"
	http "net/http"
)

type CommandHttp struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
	TTL   int64  `json:"ttl,omitempty"`
}

func handleSet(w http.ResponseWriter, r *http.Request, js *JStore) {
	var cmd CommandHttp
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := js.Execute(Command{
		Op:    "set",
		Key:   cmd.Key,
		Value: cmd.Value,
		TTL:   cmd.TTL,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleGet(w http.ResponseWriter, r *http.Request, js *JStore) {
	var cmd CommandHttp
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := js.Execute(Command{
		Op:  "get",
		Key: cmd.Key,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleDelete(w http.ResponseWriter, r *http.Request, js *JStore) {
	var cmd CommandHttp
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := js.Execute(Command{
		Op:  "delete",
		Key: cmd.Key,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func startHTTPServer(port string, js *JStore) error {
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		handleSet(w, r, js)
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		handleGet(w, r, js)
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		handleDelete(w, r, js)
	})
	return http.ListenAndServe(":"+port, nil)
}

func (js *JStore) ListenAndServeHTTP(addr string) error {
	return startHTTPServer(addr, js)
}
