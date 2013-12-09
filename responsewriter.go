package chromelogger

import (
  "fmt"
  "github.com/pilu/traffic"
)

type responseWriter struct {
  traffic.ResponseWriter
  logger  *Logger
  flushed bool
}

func (w *responseWriter) flush() {
  w.Header().Set("X-ChromeLogger-Data", w.logger.Export())
  w.flushed = true
}

func (w *responseWriter) Write(data []byte) (n int, err error) {
  if !w.flushed {
    w.flush()
  }

  return w.ResponseWriter.Write(data)
}

func (w *responseWriter) WriteHeader(statusCode int) {
  if !w.flushed {
    w.flush()
  }

  w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) WriteText(textFormat string, data ...interface{}) {
  fmt.Fprintf(w, textFormat, data...)
}
