package main

import (
     "github.com/rupak26/Real-time-Leaderboard/cmd"
     "github.com/rupak26/Real-time-Leaderboard/logger"
     "log/slog"
)

func main() {
     logs := logger.SetupLogger()
     slog.SetDefault(logs)
     slog.Info("Application Started")
     cmd.Server() 
}