package main

import (
	"yogo/pkg/context"
	"yogo/pkg/server"
)

func main() {
	ctx := context.CreateContext()
	server.CreateServerInstance(ctx)

	context.Config{}
}
