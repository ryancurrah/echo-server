package main

import (
	"fmt"
	"net/http"
	"os"

	echoserver "github.com/ryancurrah/echo-server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Echo server listening on port %s.\n", port)

	err := http.ListenAndServe(
		":"+port,
		h2c.NewHandler(
			http.HandlerFunc(echoserver.Handler),
			&http2.Server{},
		),
	)
	if err != nil {
		panic(err)
	}
}
