package chromelogger

import (
  "net/http"
  "github.com/pilu/traffic"
)

type ChromeLoggerMiddleware struct {}

func (middleware ChromeLoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next traffic.NextMiddlewareFunc) (http.ResponseWriter, *http.Request) {
  arw := w.(*traffic.AppResponseWriter)
  logger := NewLogger()
  arw.SetVar("chromelogger", logger)

  if nextMiddleware := next(); nextMiddleware != nil {
    w, r = nextMiddleware.ServeHTTP(w, r, next)
  }

  w.Header().Set("X-ChromeLogger-Data", logger.Export())

  return w, r
}

func New() *ChromeLoggerMiddleware {
  middleware := &ChromeLoggerMiddleware{}

  return middleware
}
