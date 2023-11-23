package config

import (
	"errors"
	"fmt"
	"io"
	"log/syslog"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var levels = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

type Logger struct {
	*logrus.Logger
}

func NewLogger(conf *Config) (*Logger, error) {
	c := conf.Sub("logger")

	level, ok := levels[c.GetString("level")]
	if !ok {
		return nil, errors.New("log level error")
	}

	outs := make(map[string]bool)
	for _, i := range strings.Split(c.GetString("out"), "|") {
		outs[i] = true
	}

	outWriters := make([]io.Writer, 0)
	for k := range outs {
		switch k {
		case "file":
			name := c.GetString("filename")
			if name == "" {
				return nil, errors.New("You need set one log file")
			}
			format := c.GetString("file_format")
			if format == "" {
				format = "%Y%m%d"
			}
			rotationCount := c.GetUint("file_rotation_count")
			if rotationCount == 0 {
				rotationCount = 3
			}
			opts := []rotatelogs.Option{
				rotatelogs.WithLinkName(name),
			}
			if rotationTime := c.GetInt("file_rotation_time"); rotationTime != 0 {
				opts = append(opts, rotatelogs.WithRotationTime(time.Duration(rotationTime)))
			}
			maxAge := c.GetInt("file_max_age")
			if maxAge == 0 && rotationCount == 0 {
				opts = append(opts, rotatelogs.WithMaxAge(72*time.Hour))
			} else if rotationCount != 0 {
				opts = append(opts, rotatelogs.WithRotationCount(rotationCount))
			} else if maxAge != 0 {
				opts = append(opts, rotatelogs.WithMaxAge(time.Duration(maxAge)))
			}
			out, err := rotatelogs.New(fmt.Sprintf("%s.%s", name, format), opts...)
			if err != nil {
				return nil, err
			}
			outWriters = append(outWriters, out)
		case "syslog":
			out, err := syslog.Dial("", "", syslog.LOG_WARNING|syslog.LOG_SYSLOG, "filesea")
			if err != nil {
				return nil, err
			}
			outWriters = append(outWriters, out)
		case "stdout":
			outWriters = append(outWriters, os.Stdout)
		case "stderr":
			outWriters = append(outWriters, os.Stderr)
		}
	}
	if len(outWriters) == 0 {
		outWriters = append(outWriters, os.Stdout)
	}
	timestamp := c.GetBool("timestamp")
	return &Logger{
		&logrus.Logger{
			Out: io.MultiWriter(outWriters...),
			Formatter: &logrus.TextFormatter{
				DisableTimestamp: !timestamp,
				FullTimestamp:    timestamp,
			},
			Level: level,
		},
	}, nil
}
