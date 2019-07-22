// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package promlog defines standardised ways to initialize Go kit loggers
// across Prometheus components.
// It should typically only ever be imported by main packages.
package logger

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"io"
	"os"
	stdLog "log"
	"time"
)

var (
	// This timestamp format differs from RFC3339Nano by using .000 instead
	// of .999999999 which changes the timestamp from 9 variable to 3 fixed
	// decimals (.130 instead of .130987456).
	timestampFormat = log.TimestampFormat(
		func() time.Time { return time.Now().UTC() },
		"2006-01-02T15:04:05.000Z07:00",
	)
	Logger log.Logger
)

// AllowedLevel is a settable identifier for the minimum level a log entry
// must be have.
type AllowedLevel struct {
	s string
	o level.Option
}

func (l *AllowedLevel) String() string {
	return l.s
}

// Set updates the value of the allowed level.
func (l *AllowedLevel) Set(s string) error {
	switch s {
	case "debug":
		l.o = level.AllowDebug()
	case "info":
		l.o = level.AllowInfo()
	case "warn":
		l.o = level.AllowWarn()
	case "error":
		l.o = level.AllowError()
	default:
		return errors.Errorf("unrecognized log level %q", s)
	}
	l.s = s
	return nil
}

// New returns a new leveled oklog logger. Each logged line will be annotated
// with a timestamp. The output always goes to stderr.
func InitLog(allowFormat string, allowLevel string, filename string) {
	var (
		al   = &AllowedLevel{}
		file io.Writer
	)
	al.Set(allowLevel)

	if fd, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		file = fd
	} else {
		file = os.Stderr
	}
	if allowFormat == "json" {
		Logger = log.NewJSONLogger(log.NewSyncWriter(file))
	} else {
		Logger = log.NewLogfmtLogger(log.NewSyncWriter(file))
	}

	Logger = level.NewFilter(Logger, al.o)
	Logger = log.With(Logger, "ts", timestampFormat, "caller", log.DefaultCaller)
}

func Println(msg ...interface{}) {
	level.Debug(Logger).Log(msg...)
}

func Debugln(msg ...interface{}) {
	level.Debug(Logger).Log(msg...)
}

func Infoln(msg ...interface{}) {
	level.Info(Logger).Log(msg...)
}

func Warnln(msg ...interface{}) {
	level.Warn(Logger).Log(msg...)
}

func Errorln(msg ...interface{}) {
	level.Error(Logger).Log(msg...)
}

func Fatalln(msg ...interface{}) {
	stdLog.Fatalln(msg...)
}
