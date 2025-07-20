package main

import (
	"fmt"
	"net/http"
	"otp/src/pkg/config"
	"otp/src/repo/adapter"
)

func main() {
	cnf := config.GetAppConfigInstance()
	adapter.GetReposInstances()
	fmt.Printf(fmt.Sprintf("Server running at %d", cnf.Server.Port))
	fmt.Print(http.ListenAndServe(fmt.Sprintf(":%d", cnf.Server.Port), nil))
}
