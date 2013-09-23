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

type chromeLogger struct {
  data *Data
}

func (logger *chromeLogger) add(item interface{}, logType string) {
  _, file, line, ok := runtime.Caller(2)
  var backtrace string
  if ok {
    backtrace = fmt.Sprintf("%s:%d", file, line)
  }

  logger.data.AddRow(item, backtrace, logType)
}

func (logger *chromeLogger) Log(item interface{}) {
  logger.add(item, LOG_TYPE_LOG)
}

func (logger *chromeLogger) Warn(item interface{}) {
  logger.add(item, LOG_TYPE_WARN)
}

func (logger *chromeLogger) Error(item interface{}) {
  logger.add(item, LOG_TYPE_ERROR)
}

func (logger *chromeLogger) Info(item interface{}) {
  logger.add(item, LOG_TYPE_INFO)
}

func (logger chromeLogger) Export() string {
  jsonBytes, err := json.Marshal(logger.data)
  if err != nil {
    return ""
  }

  encodedData := base64.StdEncoding.EncodeToString(jsonBytes)
  finalData := strings.Replace(encodedData, "\n", "", -1)

  return finalData
}

func newLogger() *chromeLogger {
  logger := &chromeLogger{
    data: NewData(),
  }

  return logger
}
