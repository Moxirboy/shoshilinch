package main

import (
	"shoshilinch/internal/config"
	"shoshilinch/internal/server"
	"shoshilinch/pkg/log"
)




func main() {
   var (
    config =config.Load()
   )
   log:=log.NewLogger(config.AppName,config.AppVersion)
   log.InitLogger()

   log.InfoF(
    "AppVersion: %s, LogLevel: %s, Mode: %s",
    config.AppVersion,
    config.Logger.Level,
    config.Server.Environment,
)
    server:=server.New(config,log)
    server.Run()
}
