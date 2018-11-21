package main

// main boostraper file

import (
	"yogo/pkg/di"
	"yogo/pkg/server"
)

func main() {
	container := di.CreateDependenciesContainer()
	r := server.CreateServerInstance(container)

	r.Run(container.Port())
}
