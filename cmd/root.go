package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jilleJr/avgrep/pkg/flagtype"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var config = Config{
	LogLevel:  flagtype.LogLevel(zerolog.InfoLevel),
	LogColor:  "auto",
	LogFormat: flagtype.LogFormatPretty,
}

type Config struct {
	LogLevel  flagtype.LogLevel
	LogColor  string
	LogFormat flagtype.LogFormat
}

var verbose bool
var noColorAuto = color.NoColor

var rootCmd = &cobra.Command{
	Use:           "avgrep",
	Short:         "Perform searches in ansible-vault files",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	resetLogger(zerolog.InfoLevel, flagtype.LogFormatPretty)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Var(&config.LogFormat, "log-format", "Logging output format")
	rootCmd.PersistentFlags().Var(&config.LogLevel, "log-level", "Logging severity")
}

func resetLogger(level zerolog.Level, format flagtype.LogFormat) {
	if format == flagtype.LogFormatJSON {
		resetLoggerJSON(level)
	} else {
		resetLoggerPretty(level)
	}
}

func resetLoggerPretty(level zerolog.Level) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "Jan-02 15:04",
	}).Level(level)
}

func resetLoggerJSON(level zerolog.Level) {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(level)
}

func resetColor(logColor string) {
	// supports more than documented, just to catch brain-farts
	switch strings.ToLower(strings.TrimSpace(logColor)) {
	case "auto", "":
		color.NoColor = noColorAuto
	case "never", "off", "false", "f", "0":
		color.NoColor = true
	case "always", "on", "true", "t", "1":
		color.NoColor = false
	default:
		log.Error().
			Str("log-color", config.LogColor).
			Msg("Invalid --log-color value. Must be one of: auto, never, always")
	}
}

func prettyPath(path string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}
	if home == "/" {
		// Docker containers with random USER doesn't have a home dir
		return path
	}
	withoutHome := strings.TrimPrefix(path, home)
	if withoutHome != home {
		return filepath.Join("~", withoutHome)
	}
	return path
}
