package main

import (
	"yogo/pkg/di"
	"yogo/pkg/server"
)

func main() {

	ctx := di.CreateDependenciesContainer()
	r := server.CreateServerInstance(ctx)
	r.Run(":3000")
}
