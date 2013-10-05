package chromelogger

import (
  "net/http"
  "github.com/pilu/traffic"
)

type ChromeLoggerMiddleware struct {}

func (middleware ChromeLoggerMiddleware) ServeHTTP(w traffic.ResponseWriter, r *http.Request, next traffic.NextMiddlewareFunc) (traffic.ResponseWriter, *http.Request) {
  logger := newLogger()

  flushed := false

  flush := func() {
    w.Header().Set("X-ChromeLogger-Data", logger.Export())
    flushed = true
  }

  w.SetVar("chromelogger", logger)
  w.AddBeforeWriteHandler(flush)

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
