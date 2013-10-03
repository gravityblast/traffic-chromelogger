package chromelogger

import (
  "fmt"
  "strings"
  "runtime"
  "encoding/base64"
  "encoding/json"
)

const LOG_TYPE_LOG   = "log"
const LOG_TYPE_WARN  = "warn"
const LOG_TYPE_ERROR = "error"
const LOG_TYPE_INFO  = "info"

type Logger struct {
  data *Data
}

func (logger *Logger) add(item interface{}, logType string) {
  _, file, line, ok := runtime.Caller(2)
  var backtrace string
  if ok {
    backtrace = fmt.Sprintf("%s:%d", file, line)
  }

  logger.data.AddRow(item, backtrace, logType)
}

func (logger *Logger) Log(item interface{}) {
  logger.add(item, LOG_TYPE_LOG)
}

func (logger *Logger) Warn(item interface{}) {
  logger.add(item, LOG_TYPE_WARN)
}

func (logger *Logger) Error(item interface{}) {
  logger.add(item, LOG_TYPE_ERROR)
}

func (logger *Logger) Info(item interface{}) {
  logger.add(item, LOG_TYPE_INFO)
}

func (logger Logger) Export() string {
  jsonBytes, err := json.Marshal(logger.data)
  if err != nil {
    return ""
  }

  encodedData := base64.StdEncoding.EncodeToString(jsonBytes)
  finalData := strings.Replace(encodedData, "\n", "", -1)

  return finalData
}

func newLogger() *Logger {
  logger := &Logger{
    data: NewData(),
  }

  return logger
}
