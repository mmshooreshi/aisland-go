package main

import (
	"appnest.co/aisland-go/api"
)

func main() {
	r := api.SetupRouter()

	server := startServer(r)

	handleShutdown(server)
}
