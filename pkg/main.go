package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"yogo/pkg/di"
	"yogo/pkg/server"
)

// main boostraper file

func main() {
	container := di.CreateDependenciesContainer()
	r := server.CreateServerInstance(container)

	srv := &http.Server{
		Addr:    container.Port(),
		Handler: r,
	}

	fmt.Println("Server started at :3000")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server has been interupt")
		fmt.Println(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM)
	signal.Notify(sig, os.Interrupt)
	s := <-sig
	fmt.Printf("\n Gracefully shutdown... to server (%s)\n", s.String())
	ctx := context.Background()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server has been shutdown")
	}

}
