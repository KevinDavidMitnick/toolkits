package logrus

import (
	"encoding/json"
	"runtime"

	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type MyJSONFormatter struct {
	Time  string `json:"time"`
	File  string `json:"file"`
	Line  int    `json:"line"`
	Level string `json:"level"`
	Info  string `json:"info"`
	Msg   string `json:"msg"`
}

func (f *MyJSONFormatter) Format(entry *log.Entry) ([]byte, error) {

	logrusJF := &(log.JSONFormatter{})
	bytes, _ := logrusJF.Format(entry)

	json.Unmarshal(bytes, &f)
	if _, file, no, ok := runtime.Caller(8); ok {
		f.File = file
		f.Line = no
	}

	index := strings.Index(f.Time, "+")
	times := strings.Replace(f.Time[0:index], "T", " ", 1)
	str := fmt.Sprintf("[%s] %s %s:%d %s\n", f.Level, times, f.File, f.Line, f.Msg)
	return []byte(str), nil
}

func init() {
	log.SetFormatter(&MyJSONFormatter{})
}
func InitLog(level string) (err error) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Fatal("log conf only allow [debug, info,warn,error], please check your confguire")
	}

	file, err := os.OpenFile("logs/opsultra-SwitchCollector.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Debug("Failed to log to file,using default stderr")
	}

	return
}
