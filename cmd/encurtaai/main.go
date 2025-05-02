package main

import (
	"github.com/yrnThiago/encurtador_url/config"
	"github.com/yrnThiago/encurtador_url/server"
)

func main() {
	config.Init()
	config.DatabaseInit()
	server.Init()
}
