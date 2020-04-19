package main

import (
	"github.com/liqiang/thanos/api"
	"net/http"
	"time"
)

func main() {

	routerInit := api.InitRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        routerInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
