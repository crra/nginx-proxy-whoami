// based on: https://github.com/jwilder/whoami
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	envPort     = "PORT"
	defaultPort = "8080"
)

// getEnv returns the value for an environment variable or a given default.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func main() {
	port := getEnv(envPort, defaultPort)
	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm %s\n", hostname)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
