package chromelogger

import (
  "net/http"
  "github.com/pilu/traffic"
)

type ChromeLoggerMiddleware struct {}

var Logger *chromeLogger

func (middleware ChromeLoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next traffic.NextMiddlewareFunc) (http.ResponseWriter, *http.Request) {
  Logger = newLogger()

  flushed := false

  flush := func() {
    w.Header().Set("X-ChromeLogger-Data", Logger.Export())
    flushed = true
  }

  arw := w.(*traffic.AppResponseWriter)
  arw.AddBeforeWriteHandler(flush)

  if nextMiddleware := next(); nextMiddleware != nil {
    w, r = nextMiddleware.ServeHTTP(w, r, next)
  }

  if !flushed {
    flush()
  }

  return w, r
}

func New() *ChromeLoggerMiddleware {
  middleware := &ChromeLoggerMiddleware{}

  return middleware
}
