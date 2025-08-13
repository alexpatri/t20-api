package main

import (
	"t20-api/internal/bootstrap"
)

func main() {
	flags := bootstrap.InitializeFlags()
	bootstrap.HandleMode(*flags.Env)
	cfg := bootstrap.CreateConfigContext(flags)
	bootstrap.ConnectToDatabase(cfg.Database)
	bootstrap.CreateServer(cfg)
}
