package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	myEnvVar := os.Getenv("MY_ENV_VAR")

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the pod %s, MY_ENV_VAR: %s\n", hostname, myEnvVar)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
