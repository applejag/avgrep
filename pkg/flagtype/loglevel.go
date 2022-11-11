package flagtype

import (
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
)

// LogLevel is an enum of the different supported logging severity levels.
type LogLevel zerolog.Level

func init() {
	var x LogLevel
	var _ pflag.Value = &x
}

func (f *LogLevel) String() string {
	return zerolog.Level(*f).String()
}

func (f *LogLevel) Set(value string) error {
	parsed, err := zerolog.ParseLevel(value)
	if err != nil {
		return err
	}
	*f = LogLevel(parsed)
	return nil
}

func (f *LogLevel) Type() string {
	return "level"
}
