package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("p", 80, "port for server")
	flag.Parse()
	lis := fmt.Sprintf("0.0.0.0:%d", *port)
	fmt.Println("server started")
	fmt.Printf("Please visit on your browser: http://localhost:%d\n", *port)
	if err := http.ListenAndServe(lis, http.FileServer(http.Dir("."))); err != nil {
		fmt.Printf("server start error: %v\n", err)
	}
	fmt.Println("server closed")
}
