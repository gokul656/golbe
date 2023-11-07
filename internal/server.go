package internal

import (
	"fmt"
	"net/http"

	"github.com/gokul656/golbe/config"
)

func ListenAndServe() {
	env := config.EnvConf
	addr := fmt.Sprintf(":%d", env.Port)
	server := &http.Server{
		Addr: addr,
	}

	fmt.Println("Server listening on", addr)
	http.HandleFunc("/", Forwarder)
	server.ListenAndServe()
}

func ListenAndServeV2() {
	env := config.EnvConf
	addr := fmt.Sprintf(":%d", env.Port)
	server := &http.Server{
		Addr: addr,
	}

	lb := config.LoadLBConfig().Config

	for key := range lb {
		http.HandleFunc(key, Forwarder)
	}

	fmt.Println("Server listening on", addr)
	server.ListenAndServe()
}
