package telemetry

import (
	"log/slog"
	"os"
)

var StdLog = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true,
	Level:     slog.LevelDebug,
}))

func InitLogger() {
	slog.SetDefault(StdLog)
}