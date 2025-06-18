package main

import (
	"github.com/yrnThiago/encurtador_url/config"
	"github.com/yrnThiago/encurtador_url/internal/server"
)

func main() {
	config.Init()
	config.LoggerInit()
	config.DatabaseInit()

	server.Init()
}
