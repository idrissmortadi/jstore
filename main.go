package main

import (
	js "github.com/idrissmortadi/small_go_projects/jstore/jstore"
)

func main() {
	jstore := js.NewJStore()
	// if err := jstore.ListenAndServe(":8080"); err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
	jstore.ListenAndServeHTTP("8080")
}
