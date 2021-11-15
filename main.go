package main

import (
	"juanfer2/go-queries/src/servers"
)

func main() {
	server := &servers.Server{}
	server.Start("4000")
}
