package flagtype

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// LogFormat is an enum of the different supported logging output formats.
type LogFormat string

// Enum values for the [LogFormat] type.
const (
	LogFormatPretty LogFormat = "pretty"
	LogFormatJSON   LogFormat = "json"
)

func init() {
	var x LogFormat
	var _ pflag.Value = &x
}

func (f *LogFormat) String() string {
	return string(*f)
}

func (f *LogFormat) Set(value string) error {
	parsed, err := ParseLogFormat(value)
	if err != nil {
		return err
	}
	*f = parsed
	return nil
}

func (f *LogFormat) Type() string {
	return "format"
}

func ParseLogFormat(format string) (LogFormat, error) {
	switch strings.ToLower(strings.TrimSpace(format)) {
	case "pretty":
		return LogFormatPretty, nil
	case "json":
		return LogFormatJSON, nil
	default:
		return "", fmt.Errorf(`invalid log format %q, want "pretty" or "json"`, format)
	}
}
