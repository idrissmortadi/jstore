package main

import (
	js "github.com/idrissmortadi/jstore/jstore"
)

func main() {
	ttlEnabled := true // Set to false to disable TTL
	jstore := js.NewJStore(ttlEnabled)
	// if err := jstore.ListenAndServe(":8080"); err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
	jstore.ListenAndServeHTTP("8080")
}
